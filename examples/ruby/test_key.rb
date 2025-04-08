#!/usr/bin/env ruby

require 'httparty'
require 'json'
require 'optparse'
require 'colorize'
require_relative 'lean_vault'

class OpenRouterExample
  OPENROUTER_API_URL = 'https://openrouter.ai/api/v1/chat/completions'
  
  def initialize(options = {})
    @debug = options[:debug] || false
  end

  def run
    # Load the API key into a constant (similar to how dotenv works)
    debug("Loading API key from Lean Vault...")
    LeanVault.load('ruby-app-key')
    
    test_openrouter
  rescue => e
    puts "✗ Error: #{e.message}".red
    debug("Full error: #{e.class}: #{e.backtrace.join("\n")}") if @debug
    exit 1
  end

  private

  def test_openrouter
    debug("Testing OpenRouter API connection...")
    
    response = HTTParty.post(
      OPENROUTER_API_URL,
      headers: {
        'Authorization' => "Bearer #{LeanVault::RUBY_APP_KEY}",
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

  opts.on("-d", "--debug", "Enable debug output") do
    options[:debug] = true
  end

  opts.on("-h", "--help", "Show this help message") do
    puts opts
    exit
  end
end.parse!

# Run the example
OpenRouterExample.new(options).run 