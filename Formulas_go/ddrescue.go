package main

// Ddrescue - GNU data recovery tool
// Homepage: https://www.gnu.org/software/ddrescue/ddrescue.html

import (
	"fmt"
	
	"os/exec"
)

func installDdrescue() {
	// Método 1: Descargar y extraer .tar.gz
	ddrescue_tar_url := "https://ftp.gnu.org/gnu/ddrescue/ddrescue-1.28.tar.lz"
	ddrescue_cmd_tar := exec.Command("curl", "-L", ddrescue_tar_url, "-o", "package.tar.gz")
	err := ddrescue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ddrescue_zip_url := "https://ftp.gnu.org/gnu/ddrescue/ddrescue-1.28.tar.lz"
	ddrescue_cmd_zip := exec.Command("curl", "-L", ddrescue_zip_url, "-o", "package.zip")
	err = ddrescue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ddrescue_bin_url := "https://ftp.gnu.org/gnu/ddrescue/ddrescue-1.28.tar.lz"
	ddrescue_cmd_bin := exec.Command("curl", "-L", ddrescue_bin_url, "-o", "binary.bin")
	err = ddrescue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ddrescue_src_url := "https://ftp.gnu.org/gnu/ddrescue/ddrescue-1.28.tar.lz"
	ddrescue_cmd_src := exec.Command("curl", "-L", ddrescue_src_url, "-o", "source.tar.gz")
	err = ddrescue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ddrescue_cmd_direct := exec.Command("./binary")
	err = ddrescue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
