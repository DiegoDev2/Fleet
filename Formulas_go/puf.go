package main

// Puf - Parallel URL fetcher
// Homepage: https://puf.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPuf() {
	// Método 1: Descargar y extraer .tar.gz
	puf_tar_url := "https://downloads.sourceforge.net/project/puf/puf/1.0.0/puf-1.0.0.tar.gz"
	puf_cmd_tar := exec.Command("curl", "-L", puf_tar_url, "-o", "package.tar.gz")
	err := puf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	puf_zip_url := "https://downloads.sourceforge.net/project/puf/puf/1.0.0/puf-1.0.0.zip"
	puf_cmd_zip := exec.Command("curl", "-L", puf_zip_url, "-o", "package.zip")
	err = puf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	puf_bin_url := "https://downloads.sourceforge.net/project/puf/puf/1.0.0/puf-1.0.0.bin"
	puf_cmd_bin := exec.Command("curl", "-L", puf_bin_url, "-o", "binary.bin")
	err = puf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	puf_src_url := "https://downloads.sourceforge.net/project/puf/puf/1.0.0/puf-1.0.0.src.tar.gz"
	puf_cmd_src := exec.Command("curl", "-L", puf_src_url, "-o", "source.tar.gz")
	err = puf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	puf_cmd_direct := exec.Command("./binary")
	err = puf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
