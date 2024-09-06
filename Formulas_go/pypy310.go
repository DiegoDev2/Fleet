package main

// Pypy310 - Implementation of Python 3 in Python
// Homepage: https://pypy.org/

import (
	"fmt"
	
	"os/exec"
)

func installPypy310() {
	// Método 1: Descargar y extraer .tar.gz
	pypy310_tar_url := "https://downloads.python.org/pypy/pypy3.10-v7.3.17-src.tar.bz2"
	pypy310_cmd_tar := exec.Command("curl", "-L", pypy310_tar_url, "-o", "package.tar.gz")
	err := pypy310_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pypy310_zip_url := "https://downloads.python.org/pypy/pypy3.10-v7.3.17-src.tar.bz2"
	pypy310_cmd_zip := exec.Command("curl", "-L", pypy310_zip_url, "-o", "package.zip")
	err = pypy310_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pypy310_bin_url := "https://downloads.python.org/pypy/pypy3.10-v7.3.17-src.tar.bz2"
	pypy310_cmd_bin := exec.Command("curl", "-L", pypy310_bin_url, "-o", "binary.bin")
	err = pypy310_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pypy310_src_url := "https://downloads.python.org/pypy/pypy3.10-v7.3.17-src.tar.bz2"
	pypy310_cmd_src := exec.Command("curl", "-L", pypy310_src_url, "-o", "source.tar.gz")
	err = pypy310_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pypy310_cmd_direct := exec.Command("./binary")
	err = pypy310_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: pypy")
exec.Command("latte", "install", "pypy")
	fmt.Println("Instalando dependencia: gdbm")
exec.Command("latte", "install", "gdbm")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: tcl-tk")
exec.Command("latte", "install", "tcl-tk")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
