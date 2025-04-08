#!/usr/bin/env ruby

require 'httparty'
require 'json'
require 'optparse'
require 'colorize'
require_relative 'lean_vault'

class KeyRotationExample
  def initialize(options = {})
    @debug = options[:debug] || false
    @key_name = options[:key_name] || 'ruby-app-key'
    @retries = options[:retries] || 3
  end

  def run
    debug("Starting key rotation example...")
    
    # Step 1: Load the current key and verify it works
    debug("Testing current key...")
    test_current_key
    
    # Step 2: Rotate the key
    debug("Rotating key...")
    rotate_key
    
    # Step 3: Verify the new key works
    debug("Testing new key...")
    test_current_key
    
    puts "✓ Key rotation completed successfully!".green
  rescue => e
    puts "✗ Error: #{e.message}".red
    debug("Full error: #{e.class}: #{e.backtrace.join("\n")}") if @debug
    exit 1
  end

  private

  def test_current_key
    # Load the key into a constant
    debug("Loading key from Lean Vault...")
    LeanVault.load(@key_name)
    key_const = @key_name.upcase.gsub('-', '_')
    
    # Test the key with OpenRouter API
    response = HTTParty.post(
      'https://openrouter.ai/api/v1/chat/completions',
      headers: {
        'Authorization' => "Bearer #{LeanVault.const_get(key_const)}",
        'Content-Type' => 'application/json',
        'HTTP-Referer' => 'https://github.com/spacebarlabs/lean_vault/examples/ruby',
      },
      body: {
        model: 'anthropic/claude-3-haiku',
        messages: [{ role: 'user', content: 'Say hello!' }]
      }.to_json
    )

    if response.success?
      puts "✓ API key test successful".green
      debug("Response: #{response.body}")
    else
      raise "API request failed: #{response.code} - #{response.body}"
    end
  end

  def rotate_key
    puts "Rotating API key '#{@key_name}'..."
    
    # Use lean_vault rotate command
    output = `lean_vault rotate #{@key_name} 2>&1`
    unless $?.success?
      raise "Failed to rotate key: #{output}"
    end
    
    puts "✓ Key rotated successfully".green
    
    # Remove the old constant so it can be reloaded
    if LeanVault.loaded?(@key_name)
      const_name = @key_name.upcase.gsub('-', '_')
      LeanVault.send(:remove_const, const_name)
    end
  end

  def debug(message)
    puts "DEBUG: #{message}".yellow if @debug
  end
end

# Parse command line options
options = {}
OptionParser.new do |opts|
  opts.banner = "Usage: key_rotation_example.rb [options]"

  opts.on("-d", "--debug", "Enable debug output") do
    options[:debug] = true
  end

  opts.on("-k", "--key NAME", "Key name to rotate (default: ruby-app-key)") do |name|
    options[:key_name] = name
  end

  opts.on("-r", "--retries N", Integer, "Number of retries (default: 3)") do |n|
    options[:retries] = n
  end

  opts.on("-h", "--help", "Show this help message") do
    puts opts
    exit
  end
end.parse!

# Run the example
KeyRotationExample.new(options).run 