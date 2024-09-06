package main

// Moe - Console text editor for ISO-8859 and ASCII
// Homepage: https://www.gnu.org/software/moe/moe.html

import (
	"fmt"
	
	"os/exec"
)

func installMoe() {
	// Método 1: Descargar y extraer .tar.gz
	moe_tar_url := "https://ftp.gnu.org/gnu/moe/moe-1.14.tar.lz"
	moe_cmd_tar := exec.Command("curl", "-L", moe_tar_url, "-o", "package.tar.gz")
	err := moe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	moe_zip_url := "https://ftp.gnu.org/gnu/moe/moe-1.14.tar.lz"
	moe_cmd_zip := exec.Command("curl", "-L", moe_zip_url, "-o", "package.zip")
	err = moe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	moe_bin_url := "https://ftp.gnu.org/gnu/moe/moe-1.14.tar.lz"
	moe_cmd_bin := exec.Command("curl", "-L", moe_bin_url, "-o", "binary.bin")
	err = moe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	moe_src_url := "https://ftp.gnu.org/gnu/moe/moe-1.14.tar.lz"
	moe_cmd_src := exec.Command("curl", "-L", moe_src_url, "-o", "source.tar.gz")
	err = moe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	moe_cmd_direct := exec.Command("./binary")
	err = moe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
