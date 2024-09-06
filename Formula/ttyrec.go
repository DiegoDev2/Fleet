package main

// Ttyrec - Terminal interaction recorder and player
// Homepage: http://0xcc.net/ttyrec/

import (
	"fmt"
	
	"os/exec"
)

func installTtyrec() {
	// Método 1: Descargar y extraer .tar.gz
	ttyrec_tar_url := "http://0xcc.net/ttyrec/ttyrec-1.0.8.tar.gz"
	ttyrec_cmd_tar := exec.Command("curl", "-L", ttyrec_tar_url, "-o", "package.tar.gz")
	err := ttyrec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttyrec_zip_url := "http://0xcc.net/ttyrec/ttyrec-1.0.8.zip"
	ttyrec_cmd_zip := exec.Command("curl", "-L", ttyrec_zip_url, "-o", "package.zip")
	err = ttyrec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttyrec_bin_url := "http://0xcc.net/ttyrec/ttyrec-1.0.8.bin"
	ttyrec_cmd_bin := exec.Command("curl", "-L", ttyrec_bin_url, "-o", "binary.bin")
	err = ttyrec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttyrec_src_url := "http://0xcc.net/ttyrec/ttyrec-1.0.8.src.tar.gz"
	ttyrec_cmd_src := exec.Command("curl", "-L", ttyrec_src_url, "-o", "source.tar.gz")
	err = ttyrec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttyrec_cmd_direct := exec.Command("./binary")
	err = ttyrec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
