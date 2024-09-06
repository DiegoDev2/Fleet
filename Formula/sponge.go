package main

// Sponge - Soak up standard input and write to a file
// Homepage: https://joeyh.name/code/moreutils/

import (
	"fmt"
	
	"os/exec"
)

func installSponge() {
	// Método 1: Descargar y extraer .tar.gz
	sponge_tar_url := "https://git.joeyh.name/index.cgi/moreutils.git/snapshot/moreutils-0.69.tar.gz"
	sponge_cmd_tar := exec.Command("curl", "-L", sponge_tar_url, "-o", "package.tar.gz")
	err := sponge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sponge_zip_url := "https://git.joeyh.name/index.cgi/moreutils.git/snapshot/moreutils-0.69.zip"
	sponge_cmd_zip := exec.Command("curl", "-L", sponge_zip_url, "-o", "package.zip")
	err = sponge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sponge_bin_url := "https://git.joeyh.name/index.cgi/moreutils.git/snapshot/moreutils-0.69.bin"
	sponge_cmd_bin := exec.Command("curl", "-L", sponge_bin_url, "-o", "binary.bin")
	err = sponge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sponge_src_url := "https://git.joeyh.name/index.cgi/moreutils.git/snapshot/moreutils-0.69.src.tar.gz"
	sponge_cmd_src := exec.Command("curl", "-L", sponge_src_url, "-o", "source.tar.gz")
	err = sponge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sponge_cmd_direct := exec.Command("./binary")
	err = sponge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
