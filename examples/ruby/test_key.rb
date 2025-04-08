#!/usr/bin/env ruby

require 'httparty'
require 'json'
require 'dotenv'
require 'optparse'
require 'colorize'

# Load environment variables from .env file if present
Dotenv.load

class LeanVaultExample
  OPENROUTER_API_URL = 'https://openrouter.ai/api/v1/chat/completions'
  
  def initialize(options = {})
    @key_name = options[:key_name] || ENV['LEAN_VAULT_KEY_NAME'] || 'ruby-app-key'
    @debug = options[:debug] || ENV['DEBUG'] == 'true'
    @api_key = nil
  end

  def run
    retrieve_key
    test_openrouter
  rescue => e
    puts "✗ Error: #{e.message}".red
    debug("Full error: #{e.class}: #{e.backtrace.join("\n")}") if @debug
    exit 1
  end

  private

  def retrieve_key
    debug("Retrieving key '#{@key_name}' from Lean Vault...")
    
    # Execute lean_vault command to get the key
    output = `lean_vault get #{@key_name} 2>&1`
    
    if $?.success?
      @api_key = output.strip
      puts "✓ Successfully retrieved key from Lean Vault".green
      debug("Key length: #{@api_key.length} characters")
    else
      raise "Failed to retrieve key: #{output}"
    end
  end

  def test_openrouter
    debug("Testing OpenRouter API connection...")
    
    response = HTTParty.post(
      OPENROUTER_API_URL,
      headers: {
        'Authorization' => "Bearer #{@api_key}",
        'Content-Type' => 'application/json',
        'HTTP-Referer' => 'https://github.com/spacebarlabs/lean_vault/examples/ruby',
      },
      body: {
        model: 'anthropic/claude-3-haiku',
        messages: [{ role: 'user', content: 'Say hello!' }]
      }.to_json
    )

    if response.success?
      puts "✓ OpenRouter API connection test successful".green
      content = JSON.parse(response.body)['choices'][0]['message']['content']
      puts "Response: #{content.inspect}"
    else
      raise "API request failed: #{response.code} - #{response.body}"
    end
  end

  def debug(message)
    puts "DEBUG: #{message}".yellow if @debug
  end
end

# Parse command line options
options = {}
OptionParser.new do |opts|
  opts.banner = "Usage: test_key.rb [options]"

  opts.on("-k", "--key-name NAME", "Name of the key in Lean Vault") do |name|
    options[:key_name] = name
  end

  opts.on("-d", "--debug", "Enable debug output") do
    options[:debug] = true
  end

  opts.on("-h", "--help", "Show this help message") do
    puts opts
    exit
  end
end.parse!

# Run the example
LeanVaultExample.new(options).run 