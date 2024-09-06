package main

// Ripmime - Extract attachments out of MIME encoded email packages
// Homepage: https://pldaniels.com/ripmime/

import (
	"fmt"
	
	"os/exec"
)

func installRipmime() {
	// Método 1: Descargar y extraer .tar.gz
	ripmime_tar_url := "https://pldaniels.com/ripmime/ripmime-1.4.0.10.tar.gz"
	ripmime_cmd_tar := exec.Command("curl", "-L", ripmime_tar_url, "-o", "package.tar.gz")
	err := ripmime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ripmime_zip_url := "https://pldaniels.com/ripmime/ripmime-1.4.0.10.zip"
	ripmime_cmd_zip := exec.Command("curl", "-L", ripmime_zip_url, "-o", "package.zip")
	err = ripmime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ripmime_bin_url := "https://pldaniels.com/ripmime/ripmime-1.4.0.10.bin"
	ripmime_cmd_bin := exec.Command("curl", "-L", ripmime_bin_url, "-o", "binary.bin")
	err = ripmime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ripmime_src_url := "https://pldaniels.com/ripmime/ripmime-1.4.0.10.src.tar.gz"
	ripmime_cmd_src := exec.Command("curl", "-L", ripmime_src_url, "-o", "source.tar.gz")
	err = ripmime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ripmime_cmd_direct := exec.Command("./binary")
	err = ripmime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
