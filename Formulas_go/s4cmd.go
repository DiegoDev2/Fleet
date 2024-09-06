package main

// S4cmd - Super S3 command-line tool
// Homepage: https://github.com/bloomreach/s4cmd

import (
	"fmt"
	
	"os/exec"
)

func installS4cmd() {
	// Método 1: Descargar y extraer .tar.gz
	s4cmd_tar_url := "https://files.pythonhosted.org/packages/42/b4/0061f4930958cd790098738659c1c39f8feaf688e698142435eedaa4ae34/s4cmd-2.1.0.tar.gz"
	s4cmd_cmd_tar := exec.Command("curl", "-L", s4cmd_tar_url, "-o", "package.tar.gz")
	err := s4cmd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	s4cmd_zip_url := "https://files.pythonhosted.org/packages/42/b4/0061f4930958cd790098738659c1c39f8feaf688e698142435eedaa4ae34/s4cmd-2.1.0.zip"
	s4cmd_cmd_zip := exec.Command("curl", "-L", s4cmd_zip_url, "-o", "package.zip")
	err = s4cmd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	s4cmd_bin_url := "https://files.pythonhosted.org/packages/42/b4/0061f4930958cd790098738659c1c39f8feaf688e698142435eedaa4ae34/s4cmd-2.1.0.bin"
	s4cmd_cmd_bin := exec.Command("curl", "-L", s4cmd_bin_url, "-o", "binary.bin")
	err = s4cmd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	s4cmd_src_url := "https://files.pythonhosted.org/packages/42/b4/0061f4930958cd790098738659c1c39f8feaf688e698142435eedaa4ae34/s4cmd-2.1.0.src.tar.gz"
	s4cmd_cmd_src := exec.Command("curl", "-L", s4cmd_src_url, "-o", "source.tar.gz")
	err = s4cmd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	s4cmd_cmd_direct := exec.Command("./binary")
	err = s4cmd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
