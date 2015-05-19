require 'yaml'

release_name = ARGV[0]
index_file = File.join(Dir.pwd, "releases", release_name, "index.yml")
h = YAML.load_file(index_file)
h['builds']
versions = h['builds'].values.map do |build|
  build['version'] = build['version'].to_i
end

new_version = versions.sort.last
puts new_version

