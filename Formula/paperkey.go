package main

// Paperkey - Extract just secret information out of OpenPGP secret keys
// Homepage: https://www.jabberwocky.com/software/paperkey/

import (
	"fmt"
	
	"os/exec"
)

func installPaperkey() {
	// Método 1: Descargar y extraer .tar.gz
	paperkey_tar_url := "https://www.jabberwocky.com/software/paperkey/paperkey-1.6.tar.gz"
	paperkey_cmd_tar := exec.Command("curl", "-L", paperkey_tar_url, "-o", "package.tar.gz")
	err := paperkey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	paperkey_zip_url := "https://www.jabberwocky.com/software/paperkey/paperkey-1.6.zip"
	paperkey_cmd_zip := exec.Command("curl", "-L", paperkey_zip_url, "-o", "package.zip")
	err = paperkey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	paperkey_bin_url := "https://www.jabberwocky.com/software/paperkey/paperkey-1.6.bin"
	paperkey_cmd_bin := exec.Command("curl", "-L", paperkey_bin_url, "-o", "binary.bin")
	err = paperkey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	paperkey_src_url := "https://www.jabberwocky.com/software/paperkey/paperkey-1.6.src.tar.gz"
	paperkey_cmd_src := exec.Command("curl", "-L", paperkey_src_url, "-o", "source.tar.gz")
	err = paperkey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	paperkey_cmd_direct := exec.Command("./binary")
	err = paperkey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
