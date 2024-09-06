package main

// Unrtf - RTF to other formats converter
// Homepage: https://www.gnu.org/software/unrtf/

import (
	"fmt"
	
	"os/exec"
)

func installUnrtf() {
	// Método 1: Descargar y extraer .tar.gz
	unrtf_tar_url := "https://ftp.gnu.org/gnu/unrtf/unrtf-0.21.10.tar.gz"
	unrtf_cmd_tar := exec.Command("curl", "-L", unrtf_tar_url, "-o", "package.tar.gz")
	err := unrtf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unrtf_zip_url := "https://ftp.gnu.org/gnu/unrtf/unrtf-0.21.10.zip"
	unrtf_cmd_zip := exec.Command("curl", "-L", unrtf_zip_url, "-o", "package.zip")
	err = unrtf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unrtf_bin_url := "https://ftp.gnu.org/gnu/unrtf/unrtf-0.21.10.bin"
	unrtf_cmd_bin := exec.Command("curl", "-L", unrtf_bin_url, "-o", "binary.bin")
	err = unrtf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unrtf_src_url := "https://ftp.gnu.org/gnu/unrtf/unrtf-0.21.10.src.tar.gz"
	unrtf_cmd_src := exec.Command("curl", "-L", unrtf_src_url, "-o", "source.tar.gz")
	err = unrtf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unrtf_cmd_direct := exec.Command("./binary")
	err = unrtf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
