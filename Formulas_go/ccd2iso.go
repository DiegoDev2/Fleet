package main

// Ccd2iso - Convert CloneCD images to ISO images
// Homepage: https://ccd2iso.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCcd2iso() {
	// Método 1: Descargar y extraer .tar.gz
	ccd2iso_tar_url := "https://downloads.sourceforge.net/project/ccd2iso/ccd2iso/ccd2iso-0.3/ccd2iso-0.3.tar.gz"
	ccd2iso_cmd_tar := exec.Command("curl", "-L", ccd2iso_tar_url, "-o", "package.tar.gz")
	err := ccd2iso_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ccd2iso_zip_url := "https://downloads.sourceforge.net/project/ccd2iso/ccd2iso/ccd2iso-0.3/ccd2iso-0.3.zip"
	ccd2iso_cmd_zip := exec.Command("curl", "-L", ccd2iso_zip_url, "-o", "package.zip")
	err = ccd2iso_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ccd2iso_bin_url := "https://downloads.sourceforge.net/project/ccd2iso/ccd2iso/ccd2iso-0.3/ccd2iso-0.3.bin"
	ccd2iso_cmd_bin := exec.Command("curl", "-L", ccd2iso_bin_url, "-o", "binary.bin")
	err = ccd2iso_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ccd2iso_src_url := "https://downloads.sourceforge.net/project/ccd2iso/ccd2iso/ccd2iso-0.3/ccd2iso-0.3.src.tar.gz"
	ccd2iso_cmd_src := exec.Command("curl", "-L", ccd2iso_src_url, "-o", "source.tar.gz")
	err = ccd2iso_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ccd2iso_cmd_direct := exec.Command("./binary")
	err = ccd2iso_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
