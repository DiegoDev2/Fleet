package main

// Semgrep - Easily detect and prevent bugs and anti-patterns in your codebase
// Homepage: https://semgrep.dev

import (
	"fmt"
	
	"os/exec"
)

func installSemgrep() {
	// Método 1: Descargar y extraer .tar.gz
	semgrep_tar_url := "https://github.com/semgrep/semgrep.git"
	semgrep_cmd_tar := exec.Command("curl", "-L", semgrep_tar_url, "-o", "package.tar.gz")
	err := semgrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	semgrep_zip_url := "https://github.com/semgrep/semgrep.git"
	semgrep_cmd_zip := exec.Command("curl", "-L", semgrep_zip_url, "-o", "package.zip")
	err = semgrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	semgrep_bin_url := "https://github.com/semgrep/semgrep.git"
	semgrep_cmd_bin := exec.Command("curl", "-L", semgrep_bin_url, "-o", "binary.bin")
	err = semgrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	semgrep_src_url := "https://github.com/semgrep/semgrep.git"
	semgrep_cmd_src := exec.Command("curl", "-L", semgrep_src_url, "-o", "source.tar.gz")
	err = semgrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	semgrep_cmd_direct := exec.Command("./binary")
	err = semgrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: dune")
exec.Command("latte", "install", "dune")
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
	fmt.Println("Instalando dependencia: opam")
exec.Command("latte", "install", "opam")
	fmt.Println("Instalando dependencia: pipenv")
exec.Command("latte", "install", "pipenv")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: tree-sitter")
exec.Command("latte", "install", "tree-sitter")
}
