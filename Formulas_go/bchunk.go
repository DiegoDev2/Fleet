package main

// Bchunk - Convert CD images from .bin/.cue to .iso/.cdr
// Homepage: http://he.fi/bchunk/

import (
	"fmt"
	
	"os/exec"
)

func installBchunk() {
	// Método 1: Descargar y extraer .tar.gz
	bchunk_tar_url := "http://he.fi/bchunk/bchunk-1.2.2.tar.gz"
	bchunk_cmd_tar := exec.Command("curl", "-L", bchunk_tar_url, "-o", "package.tar.gz")
	err := bchunk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bchunk_zip_url := "http://he.fi/bchunk/bchunk-1.2.2.zip"
	bchunk_cmd_zip := exec.Command("curl", "-L", bchunk_zip_url, "-o", "package.zip")
	err = bchunk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bchunk_bin_url := "http://he.fi/bchunk/bchunk-1.2.2.bin"
	bchunk_cmd_bin := exec.Command("curl", "-L", bchunk_bin_url, "-o", "binary.bin")
	err = bchunk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bchunk_src_url := "http://he.fi/bchunk/bchunk-1.2.2.src.tar.gz"
	bchunk_cmd_src := exec.Command("curl", "-L", bchunk_src_url, "-o", "source.tar.gz")
	err = bchunk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bchunk_cmd_direct := exec.Command("./binary")
	err = bchunk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
