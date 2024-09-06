package main

// Bison - Parser generator
// Homepage: https://www.gnu.org/software/bison/

import (
	"fmt"
	
	"os/exec"
)

func installBison() {
	// Método 1: Descargar y extraer .tar.gz
	bison_tar_url := "https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.xz"
	bison_cmd_tar := exec.Command("curl", "-L", bison_tar_url, "-o", "package.tar.gz")
	err := bison_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bison_zip_url := "https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.xz"
	bison_cmd_zip := exec.Command("curl", "-L", bison_zip_url, "-o", "package.zip")
	err = bison_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bison_bin_url := "https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.xz"
	bison_cmd_bin := exec.Command("curl", "-L", bison_bin_url, "-o", "binary.bin")
	err = bison_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bison_src_url := "https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.xz"
	bison_cmd_src := exec.Command("curl", "-L", bison_src_url, "-o", "source.tar.gz")
	err = bison_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bison_cmd_direct := exec.Command("./binary")
	err = bison_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
