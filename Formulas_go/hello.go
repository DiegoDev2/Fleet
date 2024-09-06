package main

// Hello - Program providing model for GNU coding standards and practices
// Homepage: https://www.gnu.org/software/hello/

import (
	"fmt"
	
	"os/exec"
)

func installHello() {
	// Método 1: Descargar y extraer .tar.gz
	hello_tar_url := "https://ftp.gnu.org/gnu/hello/hello-2.12.1.tar.gz"
	hello_cmd_tar := exec.Command("curl", "-L", hello_tar_url, "-o", "package.tar.gz")
	err := hello_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hello_zip_url := "https://ftp.gnu.org/gnu/hello/hello-2.12.1.zip"
	hello_cmd_zip := exec.Command("curl", "-L", hello_zip_url, "-o", "package.zip")
	err = hello_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hello_bin_url := "https://ftp.gnu.org/gnu/hello/hello-2.12.1.bin"
	hello_cmd_bin := exec.Command("curl", "-L", hello_bin_url, "-o", "binary.bin")
	err = hello_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hello_src_url := "https://ftp.gnu.org/gnu/hello/hello-2.12.1.src.tar.gz"
	hello_cmd_src := exec.Command("curl", "-L", hello_src_url, "-o", "source.tar.gz")
	err = hello_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hello_cmd_direct := exec.Command("./binary")
	err = hello_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
