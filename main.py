import os
import re
import subprocess

# Función para instalar una fórmula con Homebrew
def install_formula_with_brew(formula_name):
    try:
        # Verifica si la fórmula ya está instalada
        subprocess.check_call(["brew", "list", formula_name])
        print(f"{formula_name} ya está instalado.")
    except subprocess.CalledProcessError:
        # Instala la fórmula si no está instalada
        print(f"Instalando {formula_name}...")
        subprocess.check_call(["brew", "install", formula_name])

def parse_formula_rb(rb_path):
    with open(rb_path, 'r') as file:
        lines = file.readlines()

    formula_data = {
        "desc": "",
        "homepage": "",
        "url": "",
        "sha256": "",
        "dependencies": []
    }

    for line in lines:
        if line.strip().startswith("desc"):
            parts = line.split('"')
            if len(parts) > 1:
                formula_data["desc"] = parts[1]
        elif line.strip().startswith("homepage"):
            parts = line.split('"')
            if len(parts) > 1:
                formula_data["homepage"] = parts[1]
        elif line.strip().startswith("url"):
            parts = line.split('"')
            if len(parts) > 1:
                formula_data["url"] = parts[1]
        elif line.strip().startswith("sha256"):
            parts = line.split('"')
            if len(parts) > 1:
                formula_data["sha256"] = parts[1]
        elif "depends_on" in line and '"' in line:
            parts = line.split('"')
            if len(parts) > 1:
                dependency = parts[1]
                formula_data["dependencies"].append(dependency)

    return formula_data

def convert_rb_to_go(rb_filename, rb_subdir, go_subdir):
    formula_data = parse_formula_rb(rb_filename)

    dependencies_str = ', '.join([f'"{dep}"' for dep in formula_data['dependencies']])

    go_filename = os.path.join(go_subdir, os.path.basename(rb_filename).replace('.rb', '.go'))
    go_template = f"""
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// {os.path.splitext(os.path.basename(go_filename))[0]}Formula represents a formula in Go.
type {os.path.splitext(os.path.basename(go_filename))[0]}Formula struct {{
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}}

func (pkg {os.path.splitext(os.path.basename(go_filename))[0]}Formula) Print() {{
    fmt.Printf("Name: {os.path.splitext(os.path.basename(go_filename))[0]}\\n")
    fmt.Printf("Description: %s\\n", pkg.Description)
    fmt.Printf("Homepage: %s\\n", pkg.Homepage)
    fmt.Printf("URL: %s\\n", pkg.URL)
    fmt.Printf("Sha256: %s\\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\\n", pkg.Dependencies)
}}

func main() {{
    pkg := {os.path.splitext(os.path.basename(go_filename))[0]}Formula{{
        Description:  "{formula_data['desc']}",
        Homepage:     "{formula_data['homepage']}",
        URL:          "{formula_data['url']}",
        Sha256:       "{formula_data['sha256']}",
        Dependencies: []string{{{dependencies_str}}},
    }}

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {{
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {{
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }}
    }}

    if err := pkg.Install{os.path.splitext(os.path.basename(go_filename))[0]}(); err != nil {{
        log.Fatalf("Error during installation: %v", err)
    }}

    fmt.Println("Installation completed successfully.")
}}

func (pkg {os.path.splitext(os.path.basename(go_filename))[0]}Formula) Install{os.path.splitext(os.path.basename(go_filename))[0]}() error {{
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {{
        return fmt.Errorf("failed to download: %v", err)
    }}

    tarball := "{os.path.basename(formula_data['url'])}"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {{
        return fmt.Errorf("failed to extract tarball: %v", err)
    }}

    sourceDir := "{os.path.splitext(os.path.basename(formula_data['url']))[0]}"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {{
        return fmt.Errorf("failed to configure and install: %v", err)
    }}

    return nil
}}
"""

    with open(go_filename, 'w') as go_file:
        go_file.write(go_template)

    print(f"Converted {rb_filename} to {go_filename}")

def process_directory(rb_base_dir, go_base_dir):
    for subdir in os.listdir(rb_base_dir):
        rb_subdir = os.path.join(rb_base_dir, subdir)
        go_subdir = os.path.join(go_base_dir, subdir)

        if not os.path.exists(go_subdir):
            os.makedirs(go_subdir)

        for rb_file in os.listdir(rb_subdir):
            if rb_file.endswith('.rb'):
                rb_file_path = os.path.join(rb_subdir, rb_file)
                convert_rb_to_go(rb_file_path, rb_subdir, go_subdir)

if __name__ == "__main__":
    rb_base_dir = '/Users/diegodev/LattePkg/homebrew-core/Formula'
    go_base_dir = '/Users/diegodev/salida/go_files'
    process_directory(rb_base_dir, go_base_dir)
