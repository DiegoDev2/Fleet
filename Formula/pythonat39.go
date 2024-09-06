package main

// PythonAT39 - Interpreted, interactive, object-oriented programming language
// Homepage: https://www.python.org/

import (
	"fmt"
	
	"os/exec"
)

func installPythonAT39() {
	// Método 1: Descargar y extraer .tar.gz
	pythonat39_tar_url := "https://www.python.org/ftp/python/3.9.19/Python-3.9.19.tar.xz"
	pythonat39_cmd_tar := exec.Command("curl", "-L", pythonat39_tar_url, "-o", "package.tar.gz")
	err := pythonat39_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonat39_zip_url := "https://www.python.org/ftp/python/3.9.19/Python-3.9.19.tar.xz"
	pythonat39_cmd_zip := exec.Command("curl", "-L", pythonat39_zip_url, "-o", "package.zip")
	err = pythonat39_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonat39_bin_url := "https://www.python.org/ftp/python/3.9.19/Python-3.9.19.tar.xz"
	pythonat39_cmd_bin := exec.Command("curl", "-L", pythonat39_bin_url, "-o", "binary.bin")
	err = pythonat39_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonat39_src_url := "https://www.python.org/ftp/python/3.9.19/Python-3.9.19.tar.xz"
	pythonat39_cmd_src := exec.Command("curl", "-L", pythonat39_src_url, "-o", "source.tar.gz")
	err = pythonat39_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonat39_cmd_direct := exec.Command("./binary")
	err = pythonat39_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gdbm")
	exec.Command("latte", "install", "gdbm").Run()
	fmt.Println("Instalando dependencia: mpdecimal")
	exec.Command("latte", "install", "mpdecimal").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: libnsl")
	exec.Command("latte", "install", "libnsl").Run()
}
