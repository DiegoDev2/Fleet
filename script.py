import os
import re

# Rutas de las fórmulas Ruby y destino para las fórmulas Go
RUBY_FORMULA_PATH = "./Formula"  # Carpeta raíz donde están las fórmulas Ruby
GO_FORMULA_PATH = "./Formulas_go"  # Carpeta raíz donde guardarás las fórmulas Go

# Crear las carpetas destino si no existen
os.makedirs(GO_FORMULA_PATH, exist_ok=True)

# Expresiones regulares para extraer la información relevante de las fórmulas Ruby
regex_name = re.compile(r'class (\w+) < Formula')
regex_desc = re.compile(r'desc "(.+?)"')
regex_homepage = re.compile(r'homepage "(.+?)"')
regex_url = re.compile(r'url "(.+?)"')
regex_sha256 = re.compile(r'sha256 "(.+?)"')
regex_dependencies = re.compile(r'depends_on "(.+?)"')

def process_formula(file_path):
    """ Procesa una fórmula Ruby y extrae la información relevante """
    with open(file_path, 'r') as f:
        content = f.read()

    # Extraer el nombre, descripción, homepage, URL, SHA256 y dependencias
    name_match = regex_name.search(content)
    desc_match = regex_desc.search(content)
    homepage_match = regex_homepage.search(content)
    url_match = regex_url.search(content)
    sha256_match = regex_sha256.search(content)
    dependencies = regex_dependencies.findall(content)

    # Validar que se haya encontrado la información mínima
    if not (name_match and url_match and sha256_match):
        print(f"Error: No se pudo procesar la fórmula {file_path}")
        return None

    # Crear un diccionario con la información relevante
    formula_data = {
        "name": name_match.group(1),
        "desc": desc_match.group(1) if desc_match else "No description",
        "homepage": homepage_match.group(1) if homepage_match else "No homepage",
        "url": url_match.group(1),
        "sha256": sha256_match.group(1),
        "dependencies": dependencies
    }
    return formula_data

def generate_unique_variable(name, suffix):
    """ Genera nombres de variables únicos usando el nombre del paquete y un sufijo """
    return f"{name.lower()}_{suffix}"

def write_go_formula(formula_data, dest_path):
    """ Escribe un archivo .go con la información extraída """
    os.makedirs(os.path.dirname(dest_path), exist_ok=True)  # Crear carpeta si no existe
    with open(dest_path, 'w') as go_file:
        # Nombres de variables únicos
        tar_url_var = generate_unique_variable(formula_data["name"], "tar_url")
        zip_url_var = generate_unique_variable(formula_data["name"], "zip_url")
        bin_url_var = generate_unique_variable(formula_data["name"], "bin_url")
        src_url_var = generate_unique_variable(formula_data["name"], "src_url")
        sha256_var = generate_unique_variable(formula_data["name"], "sha256")
        cmd_var = generate_unique_variable(formula_data["name"], "cmd")

        go_file.write(f"package main\n\n")
        go_file.write(f'// {formula_data["name"]} - {formula_data["desc"]}\n')
        go_file.write(f'// Homepage: {formula_data["homepage"]}\n\n')
        go_file.write(f'import (\n\t"fmt"\n\t\n\t"os/exec"\n)\n\n')

        go_file.write(f'func install{formula_data["name"]}() {{\n')

        # Método de instalación 1: .tar.gz
        go_file.write(f'\t// Método 1: Descargar y extraer .tar.gz\n')
        go_file.write(f'\t{tar_url_var} := "{formula_data["url"]}"\n')
        go_file.write(f'\t{cmd_var}_tar := exec.Command("curl", "-L", {tar_url_var}, "-o", "package.tar.gz")\n')
        go_file.write(f'\terr := {cmd_var}_tar.Run()\n')
        go_file.write(f'\tif err != nil {{\n')
        go_file.write(f'\t\tfmt.Println("Error al descargar .tar.gz:", err)\n')
        go_file.write(f'\t\treturn\n\t}}\n')
        go_file.write(f'\texec.Command("tar", "-xzf", "package.tar.gz").Run()\n\n')

        # Método de instalación 2: .zip
        go_file.write(f'\t// Método 2: Descargar y extraer .zip\n')
        go_file.write(f'\t{zip_url_var} := "{formula_data["url"].replace(".tar.gz", ".zip")}"\n')
        go_file.write(f'\t{cmd_var}_zip := exec.Command("curl", "-L", {zip_url_var}, "-o", "package.zip")\n')
        go_file.write(f'\terr = {cmd_var}_zip.Run()\n')
        go_file.write(f'\tif err != nil {{\n')
        go_file.write(f'\t\tfmt.Println("Error al descargar .zip:", err)\n')
        go_file.write(f'\t\treturn\n\t}}\n')
        go_file.write(f'\texec.Command("unzip", "package.zip").Run()\n\n')

        # Método de instalación 3: Binario precompilado
        go_file.write(f'\t// Método 3: Descargar binario precompilado\n')
        go_file.write(f'\t{bin_url_var} := "{formula_data["url"].replace(".tar.gz", ".bin")}"\n')
        go_file.write(f'\t{cmd_var}_bin := exec.Command("curl", "-L", {bin_url_var}, "-o", "binary.bin")\n')
        go_file.write(f'\terr = {cmd_var}_bin.Run()\n')
        go_file.write(f'\tif err != nil {{\n')
        go_file.write(f'\t\tfmt.Println("Error al descargar binario:", err)\n')
        go_file.write(f'\t\treturn\n\t}}\n')
        go_file.write(f'\texec.Command("chmod", "+x", "binary.bin").Run()\n')
        go_file.write(f'\texec.Command("./binary.bin").Run()\n\n')

        # Método de instalación 4: Compilación desde código fuente
        go_file.write(f'\t// Método 4: Descargar y compilar desde código fuente\n')
        go_file.write(f'\t{src_url_var} := "{formula_data["url"].replace(".tar.gz", ".src.tar.gz")}"\n')
        go_file.write(f'\t{cmd_var}_src := exec.Command("curl", "-L", {src_url_var}, "-o", "source.tar.gz")\n')
        go_file.write(f'\terr = {cmd_var}_src.Run()\n')
        go_file.write(f'\tif err != nil {{\n')
        go_file.write(f'\t\tfmt.Println("Error al descargar código fuente:", err)\n')
        go_file.write(f'\t\treturn\n\t}}\n')
        go_file.write(f'\texec.Command("tar", "-xzf", "source.tar.gz").Run()\n')
        go_file.write(f'\texec.Command("make").Run()\n\n')

        # Método de instalación 5: Binarios directos
        go_file.write(f'\t// Método 5: Ejecutar binario directo\n')
        go_file.write(f'\t{cmd_var}_direct := exec.Command("./binary")\n')
        go_file.write(f'\terr = {cmd_var}_direct.Run()\n')
        go_file.write(f'\tif err != nil {{\n')
        go_file.write(f'\t\tfmt.Println("Error al ejecutar binario:", err)\n')
        go_file.write(f'\t\treturn\n\t}}\n')

        # Dependencias (si existen)
        if formula_data["dependencies"]:
            go_file.write(f'\t// Instalar dependencias\n')
            for dep in formula_data["dependencies"]:
                go_file.write(f'\tfmt.Println("Instalando dependencia: {dep}")\n')
                go_file.write(f'\texec.Command("latte", "install", "{dep}").Run()\n')

        go_file.write(f'}}\n')

def main():
    """ Procesa todas las fórmulas Ruby y genera las equivalentes en Go """
    for root, _, files in os.walk(RUBY_FORMULA_PATH):
        for file_name in files:
            if file_name.endswith(".rb"):
                file_path = os.path.join(root, file_name)
                formula_data = process_formula(file_path)
                if formula_data:
                    go_file_name = f"{formula_data['name'].lower()}.go"
                    dest_path = os.path.join(GO_FORMULA_PATH, go_file_name)
                    write_go_formula(formula_data, dest_path)
                    print(f"Generada fórmula Go: {go_file_name}")

if __name__ == "__main__":
    main()
