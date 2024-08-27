import os
import re

ruby_dir = "/Users/diegodev/LattePkg/Formula/a/formulas"
go_dir = "/Users/diegodev/LattePkg/Formula/a"

url_pattern = re.compile(r'url "(.*?)"')
desc_pattern = re.compile(r'desc "(.*?)"')
homepage_pattern = re.compile(r'homepage "(.*?)"')
depends_pattern = re.compile(r'depends_on "(.*?)"')

def extract_info_from_rb(file_content):
    url = re.findall(url_pattern, file_content)
    desc = re.findall(desc_pattern, file_content)
    homepage = re.findall(homepage_pattern, file_content)
    depends = re.findall(depends_pattern, file_content)

    return {
        "url": url[0] if url else "",
        "description": desc[0] if desc else "",
        "homepage": homepage[0] if homepage else "",
        "dependencies": depends if depends else []
    }

def update_go_file(go_file_path, info, tool_name):
    try:
        with open(go_file_path, 'r') as file:
            content = file.read()

        # Prefijar con el nombre de la herramienta para evitar conflictos
        content = re.sub(r'(Description\s*=\s*")[^"]*', r'\1' + info['description'], content)
        content = re.sub(r'(Homepage\s*=\s*")[^"]*', r'\1' + info['homepage'], content)
        content = re.sub(r'(URL\s*=\s*")[^"]*', r'\1' + info['url'], content)

        dependencies = ', '.join(['"{}"'.format(dep) for dep in info['dependencies']])
        content = re.sub(r'(Dependencies\s*=\s*\[)[^\]]*', r'\1' + dependencies, content)

        # Guardar el contenido actualizado
        with open(go_file_path, 'w') as file:
            file.write(content)

        print(f"Actualizado {go_file_path} con éxito.")
    except Exception as e:
        print(f"Error al actualizar {go_file_path}: {e}")

def main():
    for ruby_file in os.listdir(ruby_dir):
        if ruby_file.endswith(".rb"):
            base_name = ruby_file[:-3]  # Remove .rb extension
            go_file = f"{base_name}.go"
            go_file_path = os.path.join(go_dir, go_file)

            if os.path.exists(go_file_path):
                ruby_file_path = os.path.join(ruby_dir, ruby_file)

                with open(ruby_file_path, 'r') as file:
                    rb_content = file.read()

                info = extract_info_from_rb(rb_content)

                # Pasar el nombre base para usarlo como prefijo y evitar conflictos
                update_go_file(go_file_path, info, base_name)
            else:
                print(f"Archivo Go no encontrado para {ruby_file}")

if __name__ == "__main__":
    main()
