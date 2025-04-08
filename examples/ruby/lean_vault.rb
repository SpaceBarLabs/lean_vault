# LeanVault module provides functionality similar to dotenv
# It loads secrets from lean_vault into constants at runtime
module LeanVault
  class Error < StandardError; end
  
  class << self
    def load(*keys)
      keys.each do |key_name|
        const_name = key_name.upcase.gsub('-', '_')
        value = `lean_vault get #{key_name} 2>&1`.strip
        
        if $?.success?
          const_set(const_name, value)
        else
          raise Error, "Failed to load key '#{key_name}': #{value}"
        end
      end
    end

    # Load all keys from the vault
    def load_all
      output = `lean_vault list 2>&1`
      raise Error, "Failed to list keys: #{output}" unless $?.success?
      
      keys = output.split("\n")
      load(*keys) unless keys.empty?
    end

    # Check if a key is loaded as a constant
    def loaded?(key_name)
      const_name = key_name.upcase.gsub('-', '_')
      const_defined?(const_name)
    end
  end
end 