package main

// PythonAT311 - Interpreted, interactive, object-oriented programming language
// Homepage: https://www.python.org/

import (
	"fmt"
	
	"os/exec"
)

func installPythonAT311() {
	// Método 1: Descargar y extraer .tar.gz
	pythonat311_tar_url := "https://www.python.org/ftp/python/3.11.9/Python-3.11.9.tgz"
	pythonat311_cmd_tar := exec.Command("curl", "-L", pythonat311_tar_url, "-o", "package.tar.gz")
	err := pythonat311_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonat311_zip_url := "https://www.python.org/ftp/python/3.11.9/Python-3.11.9.tgz"
	pythonat311_cmd_zip := exec.Command("curl", "-L", pythonat311_zip_url, "-o", "package.zip")
	err = pythonat311_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonat311_bin_url := "https://www.python.org/ftp/python/3.11.9/Python-3.11.9.tgz"
	pythonat311_cmd_bin := exec.Command("curl", "-L", pythonat311_bin_url, "-o", "binary.bin")
	err = pythonat311_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonat311_src_url := "https://www.python.org/ftp/python/3.11.9/Python-3.11.9.tgz"
	pythonat311_cmd_src := exec.Command("curl", "-L", pythonat311_src_url, "-o", "source.tar.gz")
	err = pythonat311_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonat311_cmd_direct := exec.Command("./binary")
	err = pythonat311_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: mpdecimal")
exec.Command("latte", "install", "mpdecimal")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: berkeley-db@5")
exec.Command("latte", "install", "berkeley-db@5")
	fmt.Println("Instalando dependencia: libnsl")
exec.Command("latte", "install", "libnsl")
}
