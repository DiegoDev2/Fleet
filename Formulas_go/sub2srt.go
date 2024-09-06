package main

// Sub2srt - Convert subtitles from .sub to subviewer .srt format
// Homepage: https://github.com/robelix/sub2srt

import (
	"fmt"
	
	"os/exec"
)

func installSub2srt() {
	// Método 1: Descargar y extraer .tar.gz
	sub2srt_tar_url := "https://github.com/robelix/sub2srt/archive/refs/tags/0.5.5.tar.gz"
	sub2srt_cmd_tar := exec.Command("curl", "-L", sub2srt_tar_url, "-o", "package.tar.gz")
	err := sub2srt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sub2srt_zip_url := "https://github.com/robelix/sub2srt/archive/refs/tags/0.5.5.zip"
	sub2srt_cmd_zip := exec.Command("curl", "-L", sub2srt_zip_url, "-o", "package.zip")
	err = sub2srt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sub2srt_bin_url := "https://github.com/robelix/sub2srt/archive/refs/tags/0.5.5.bin"
	sub2srt_cmd_bin := exec.Command("curl", "-L", sub2srt_bin_url, "-o", "binary.bin")
	err = sub2srt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sub2srt_src_url := "https://github.com/robelix/sub2srt/archive/refs/tags/0.5.5.src.tar.gz"
	sub2srt_cmd_src := exec.Command("curl", "-L", sub2srt_src_url, "-o", "source.tar.gz")
	err = sub2srt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sub2srt_cmd_direct := exec.Command("./binary")
	err = sub2srt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
