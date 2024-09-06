package main

// BpmTools - Detect tempo of audio files using beats-per-minute (BPM)
// Homepage: https://www.pogo.org.uk/~mark/bpm-tools/

import (
	"fmt"
	
	"os/exec"
)

func installBpmTools() {
	// Método 1: Descargar y extraer .tar.gz
	bpmtools_tar_url := "https://www.pogo.org.uk/~mark/bpm-tools/releases/bpm-tools-0.3.tar.gz"
	bpmtools_cmd_tar := exec.Command("curl", "-L", bpmtools_tar_url, "-o", "package.tar.gz")
	err := bpmtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bpmtools_zip_url := "https://www.pogo.org.uk/~mark/bpm-tools/releases/bpm-tools-0.3.zip"
	bpmtools_cmd_zip := exec.Command("curl", "-L", bpmtools_zip_url, "-o", "package.zip")
	err = bpmtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bpmtools_bin_url := "https://www.pogo.org.uk/~mark/bpm-tools/releases/bpm-tools-0.3.bin"
	bpmtools_cmd_bin := exec.Command("curl", "-L", bpmtools_bin_url, "-o", "binary.bin")
	err = bpmtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bpmtools_src_url := "https://www.pogo.org.uk/~mark/bpm-tools/releases/bpm-tools-0.3.src.tar.gz"
	bpmtools_cmd_src := exec.Command("curl", "-L", bpmtools_src_url, "-o", "source.tar.gz")
	err = bpmtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bpmtools_cmd_direct := exec.Command("./binary")
	err = bpmtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
