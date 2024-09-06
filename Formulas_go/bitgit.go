package main

// BitGit - Bit is a modern Git CLI
// Homepage: https://github.com/chriswalz/bit

import (
	"fmt"
	
	"os/exec"
)

func installBitGit() {
	// Método 1: Descargar y extraer .tar.gz
	bitgit_tar_url := "https://github.com/chriswalz/bit/archive/refs/tags/v1.1.2.tar.gz"
	bitgit_cmd_tar := exec.Command("curl", "-L", bitgit_tar_url, "-o", "package.tar.gz")
	err := bitgit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bitgit_zip_url := "https://github.com/chriswalz/bit/archive/refs/tags/v1.1.2.zip"
	bitgit_cmd_zip := exec.Command("curl", "-L", bitgit_zip_url, "-o", "package.zip")
	err = bitgit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bitgit_bin_url := "https://github.com/chriswalz/bit/archive/refs/tags/v1.1.2.bin"
	bitgit_cmd_bin := exec.Command("curl", "-L", bitgit_bin_url, "-o", "binary.bin")
	err = bitgit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bitgit_src_url := "https://github.com/chriswalz/bit/archive/refs/tags/v1.1.2.src.tar.gz"
	bitgit_cmd_src := exec.Command("curl", "-L", bitgit_src_url, "-o", "source.tar.gz")
	err = bitgit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bitgit_cmd_direct := exec.Command("./binary")
	err = bitgit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
