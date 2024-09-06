package main

// Nrg2iso - Extract ISO9660 data from Nero nrg files
// Homepage: http://gregory.kokanosky.free.fr/v4/linux/nrg2iso.en.html

import (
	"fmt"
	
	"os/exec"
)

func installNrg2iso() {
	// Método 1: Descargar y extraer .tar.gz
	nrg2iso_tar_url := "http://gregory.kokanosky.free.fr/v4/linux/nrg2iso-0.4.1.tar.gz"
	nrg2iso_cmd_tar := exec.Command("curl", "-L", nrg2iso_tar_url, "-o", "package.tar.gz")
	err := nrg2iso_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nrg2iso_zip_url := "http://gregory.kokanosky.free.fr/v4/linux/nrg2iso-0.4.1.zip"
	nrg2iso_cmd_zip := exec.Command("curl", "-L", nrg2iso_zip_url, "-o", "package.zip")
	err = nrg2iso_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nrg2iso_bin_url := "http://gregory.kokanosky.free.fr/v4/linux/nrg2iso-0.4.1.bin"
	nrg2iso_cmd_bin := exec.Command("curl", "-L", nrg2iso_bin_url, "-o", "binary.bin")
	err = nrg2iso_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nrg2iso_src_url := "http://gregory.kokanosky.free.fr/v4/linux/nrg2iso-0.4.1.src.tar.gz"
	nrg2iso_cmd_src := exec.Command("curl", "-L", nrg2iso_src_url, "-o", "source.tar.gz")
	err = nrg2iso_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nrg2iso_cmd_direct := exec.Command("./binary")
	err = nrg2iso_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
