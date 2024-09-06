package main

// Bc - Arbitrary precision numeric processing language
// Homepage: https://www.gnu.org/software/bc/

import (
	"fmt"
	
	"os/exec"
)

func installBc() {
	// Método 1: Descargar y extraer .tar.gz
	bc_tar_url := "https://ftp.gnu.org/gnu/bc/bc-1.07.1.tar.gz"
	bc_cmd_tar := exec.Command("curl", "-L", bc_tar_url, "-o", "package.tar.gz")
	err := bc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bc_zip_url := "https://ftp.gnu.org/gnu/bc/bc-1.07.1.zip"
	bc_cmd_zip := exec.Command("curl", "-L", bc_zip_url, "-o", "package.zip")
	err = bc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bc_bin_url := "https://ftp.gnu.org/gnu/bc/bc-1.07.1.bin"
	bc_cmd_bin := exec.Command("curl", "-L", bc_bin_url, "-o", "binary.bin")
	err = bc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bc_src_url := "https://ftp.gnu.org/gnu/bc/bc-1.07.1.src.tar.gz"
	bc_cmd_src := exec.Command("curl", "-L", bc_src_url, "-o", "source.tar.gz")
	err = bc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bc_cmd_direct := exec.Command("./binary")
	err = bc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
