package main

// Esniper - Snipe eBay auctions from the command-line
// Homepage: https://sourceforge.net/projects/esniper/

import (
	"fmt"
	
	"os/exec"
)

func installEsniper() {
	// Método 1: Descargar y extraer .tar.gz
	esniper_tar_url := "https://downloads.sourceforge.net/project/esniper/esniper/2.35.0/esniper-2-35-0.tgz"
	esniper_cmd_tar := exec.Command("curl", "-L", esniper_tar_url, "-o", "package.tar.gz")
	err := esniper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	esniper_zip_url := "https://downloads.sourceforge.net/project/esniper/esniper/2.35.0/esniper-2-35-0.tgz"
	esniper_cmd_zip := exec.Command("curl", "-L", esniper_zip_url, "-o", "package.zip")
	err = esniper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	esniper_bin_url := "https://downloads.sourceforge.net/project/esniper/esniper/2.35.0/esniper-2-35-0.tgz"
	esniper_cmd_bin := exec.Command("curl", "-L", esniper_bin_url, "-o", "binary.bin")
	err = esniper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	esniper_src_url := "https://downloads.sourceforge.net/project/esniper/esniper/2.35.0/esniper-2-35-0.tgz"
	esniper_cmd_src := exec.Command("curl", "-L", esniper_src_url, "-o", "source.tar.gz")
	err = esniper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	esniper_cmd_direct := exec.Command("./binary")
	err = esniper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
