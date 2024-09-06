package main

// Mjpegtools - Record and playback videos and perform simple edits
// Homepage: https://mjpeg.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installMjpegtools() {
	// Método 1: Descargar y extraer .tar.gz
	mjpegtools_tar_url := "https://downloads.sourceforge.net/project/mjpeg/mjpegtools/2.2.1/mjpegtools-2.2.1.tar.gz"
	mjpegtools_cmd_tar := exec.Command("curl", "-L", mjpegtools_tar_url, "-o", "package.tar.gz")
	err := mjpegtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mjpegtools_zip_url := "https://downloads.sourceforge.net/project/mjpeg/mjpegtools/2.2.1/mjpegtools-2.2.1.zip"
	mjpegtools_cmd_zip := exec.Command("curl", "-L", mjpegtools_zip_url, "-o", "package.zip")
	err = mjpegtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mjpegtools_bin_url := "https://downloads.sourceforge.net/project/mjpeg/mjpegtools/2.2.1/mjpegtools-2.2.1.bin"
	mjpegtools_cmd_bin := exec.Command("curl", "-L", mjpegtools_bin_url, "-o", "binary.bin")
	err = mjpegtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mjpegtools_src_url := "https://downloads.sourceforge.net/project/mjpeg/mjpegtools/2.2.1/mjpegtools-2.2.1.src.tar.gz"
	mjpegtools_cmd_src := exec.Command("curl", "-L", mjpegtools_src_url, "-o", "source.tar.gz")
	err = mjpegtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mjpegtools_cmd_direct := exec.Command("./binary")
	err = mjpegtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
}
