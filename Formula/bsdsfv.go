package main

// Bsdsfv - SFV utility tools
// Homepage: https://bsdsfv.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBsdsfv() {
	// Método 1: Descargar y extraer .tar.gz
	bsdsfv_tar_url := "https://downloads.sourceforge.net/project/bsdsfv/bsdsfv/1.18/bsdsfv-1.18.tar.gz"
	bsdsfv_cmd_tar := exec.Command("curl", "-L", bsdsfv_tar_url, "-o", "package.tar.gz")
	err := bsdsfv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bsdsfv_zip_url := "https://downloads.sourceforge.net/project/bsdsfv/bsdsfv/1.18/bsdsfv-1.18.zip"
	bsdsfv_cmd_zip := exec.Command("curl", "-L", bsdsfv_zip_url, "-o", "package.zip")
	err = bsdsfv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bsdsfv_bin_url := "https://downloads.sourceforge.net/project/bsdsfv/bsdsfv/1.18/bsdsfv-1.18.bin"
	bsdsfv_cmd_bin := exec.Command("curl", "-L", bsdsfv_bin_url, "-o", "binary.bin")
	err = bsdsfv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bsdsfv_src_url := "https://downloads.sourceforge.net/project/bsdsfv/bsdsfv/1.18/bsdsfv-1.18.src.tar.gz"
	bsdsfv_cmd_src := exec.Command("curl", "-L", bsdsfv_src_url, "-o", "source.tar.gz")
	err = bsdsfv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bsdsfv_cmd_direct := exec.Command("./binary")
	err = bsdsfv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
