package main

// Io - Small prototype-based programming language
// Homepage: http://iolanguage.com/

import (
	"fmt"
	
	"os/exec"
)

func installIo() {
	// Método 1: Descargar y extraer .tar.gz
	io_tar_url := "https://github.com/IoLanguage/io/archive/refs/tags/2017.09.06.tar.gz"
	io_cmd_tar := exec.Command("curl", "-L", io_tar_url, "-o", "package.tar.gz")
	err := io_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	io_zip_url := "https://github.com/IoLanguage/io/archive/refs/tags/2017.09.06.zip"
	io_cmd_zip := exec.Command("curl", "-L", io_zip_url, "-o", "package.zip")
	err = io_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	io_bin_url := "https://github.com/IoLanguage/io/archive/refs/tags/2017.09.06.bin"
	io_cmd_bin := exec.Command("curl", "-L", io_bin_url, "-o", "binary.bin")
	err = io_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	io_src_url := "https://github.com/IoLanguage/io/archive/refs/tags/2017.09.06.src.tar.gz"
	io_cmd_src := exec.Command("curl", "-L", io_src_url, "-o", "source.tar.gz")
	err = io_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	io_cmd_direct := exec.Command("./binary")
	err = io_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
