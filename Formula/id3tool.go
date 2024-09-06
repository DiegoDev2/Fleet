package main

// Id3tool - ID3 editing tool
// Homepage: http://nekohako.xware.cx/id3tool/

import (
	"fmt"
	
	"os/exec"
)

func installId3tool() {
	// Método 1: Descargar y extraer .tar.gz
	id3tool_tar_url := "http://nekohako.xware.cx/id3tool/id3tool-1.2a.tar.gz"
	id3tool_cmd_tar := exec.Command("curl", "-L", id3tool_tar_url, "-o", "package.tar.gz")
	err := id3tool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	id3tool_zip_url := "http://nekohako.xware.cx/id3tool/id3tool-1.2a.zip"
	id3tool_cmd_zip := exec.Command("curl", "-L", id3tool_zip_url, "-o", "package.zip")
	err = id3tool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	id3tool_bin_url := "http://nekohako.xware.cx/id3tool/id3tool-1.2a.bin"
	id3tool_cmd_bin := exec.Command("curl", "-L", id3tool_bin_url, "-o", "binary.bin")
	err = id3tool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	id3tool_src_url := "http://nekohako.xware.cx/id3tool/id3tool-1.2a.src.tar.gz"
	id3tool_cmd_src := exec.Command("curl", "-L", id3tool_src_url, "-o", "source.tar.gz")
	err = id3tool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	id3tool_cmd_direct := exec.Command("./binary")
	err = id3tool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
