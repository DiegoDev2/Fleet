package main

// H264bitstream - Library for reading and writing H264 video streams
// Homepage: https://h264bitstream.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installH264bitstream() {
	// Método 1: Descargar y extraer .tar.gz
	h264bitstream_tar_url := "https://downloads.sourceforge.net/project/h264bitstream/h264bitstream/0.2.0/h264bitstream-0.2.0.tar.gz"
	h264bitstream_cmd_tar := exec.Command("curl", "-L", h264bitstream_tar_url, "-o", "package.tar.gz")
	err := h264bitstream_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	h264bitstream_zip_url := "https://downloads.sourceforge.net/project/h264bitstream/h264bitstream/0.2.0/h264bitstream-0.2.0.zip"
	h264bitstream_cmd_zip := exec.Command("curl", "-L", h264bitstream_zip_url, "-o", "package.zip")
	err = h264bitstream_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	h264bitstream_bin_url := "https://downloads.sourceforge.net/project/h264bitstream/h264bitstream/0.2.0/h264bitstream-0.2.0.bin"
	h264bitstream_cmd_bin := exec.Command("curl", "-L", h264bitstream_bin_url, "-o", "binary.bin")
	err = h264bitstream_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	h264bitstream_src_url := "https://downloads.sourceforge.net/project/h264bitstream/h264bitstream/0.2.0/h264bitstream-0.2.0.src.tar.gz"
	h264bitstream_cmd_src := exec.Command("curl", "-L", h264bitstream_src_url, "-o", "source.tar.gz")
	err = h264bitstream_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	h264bitstream_cmd_direct := exec.Command("./binary")
	err = h264bitstream_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
