package main

// Mpegdemux - MPEG1/2 system stream demultiplexer
// Homepage: http://www.hampa.ch/mpegdemux/

import (
	"fmt"
	
	"os/exec"
)

func installMpegdemux() {
	// Método 1: Descargar y extraer .tar.gz
	mpegdemux_tar_url := "http://www.hampa.ch/mpegdemux/mpegdemux-0.1.5.tar.gz"
	mpegdemux_cmd_tar := exec.Command("curl", "-L", mpegdemux_tar_url, "-o", "package.tar.gz")
	err := mpegdemux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpegdemux_zip_url := "http://www.hampa.ch/mpegdemux/mpegdemux-0.1.5.zip"
	mpegdemux_cmd_zip := exec.Command("curl", "-L", mpegdemux_zip_url, "-o", "package.zip")
	err = mpegdemux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpegdemux_bin_url := "http://www.hampa.ch/mpegdemux/mpegdemux-0.1.5.bin"
	mpegdemux_cmd_bin := exec.Command("curl", "-L", mpegdemux_bin_url, "-o", "binary.bin")
	err = mpegdemux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpegdemux_src_url := "http://www.hampa.ch/mpegdemux/mpegdemux-0.1.5.src.tar.gz"
	mpegdemux_cmd_src := exec.Command("curl", "-L", mpegdemux_src_url, "-o", "source.tar.gz")
	err = mpegdemux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpegdemux_cmd_direct := exec.Command("./binary")
	err = mpegdemux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
