package main

// Libffcall - GNU Foreign Function Interface library
// Homepage: https://www.gnu.org/software/libffcall/

import (
	"fmt"
	
	"os/exec"
)

func installLibffcall() {
	// Método 1: Descargar y extraer .tar.gz
	libffcall_tar_url := "https://ftp.gnu.org/gnu/libffcall/libffcall-2.5.tar.gz"
	libffcall_cmd_tar := exec.Command("curl", "-L", libffcall_tar_url, "-o", "package.tar.gz")
	err := libffcall_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libffcall_zip_url := "https://ftp.gnu.org/gnu/libffcall/libffcall-2.5.zip"
	libffcall_cmd_zip := exec.Command("curl", "-L", libffcall_zip_url, "-o", "package.zip")
	err = libffcall_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libffcall_bin_url := "https://ftp.gnu.org/gnu/libffcall/libffcall-2.5.bin"
	libffcall_cmd_bin := exec.Command("curl", "-L", libffcall_bin_url, "-o", "binary.bin")
	err = libffcall_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libffcall_src_url := "https://ftp.gnu.org/gnu/libffcall/libffcall-2.5.src.tar.gz"
	libffcall_cmd_src := exec.Command("curl", "-L", libffcall_src_url, "-o", "source.tar.gz")
	err = libffcall_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libffcall_cmd_direct := exec.Command("./binary")
	err = libffcall_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
