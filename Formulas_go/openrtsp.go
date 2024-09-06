package main

// Openrtsp - Command-line RTSP client
// Homepage: http://www.live555.com/openRTSP

import (
	"fmt"
	
	"os/exec"
)

func installOpenrtsp() {
	// Método 1: Descargar y extraer .tar.gz
	openrtsp_tar_url := "http://www.live555.com/liveMedia/public/live.2024.08.01.tar.gz"
	openrtsp_cmd_tar := exec.Command("curl", "-L", openrtsp_tar_url, "-o", "package.tar.gz")
	err := openrtsp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openrtsp_zip_url := "http://www.live555.com/liveMedia/public/live.2024.08.01.zip"
	openrtsp_cmd_zip := exec.Command("curl", "-L", openrtsp_zip_url, "-o", "package.zip")
	err = openrtsp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openrtsp_bin_url := "http://www.live555.com/liveMedia/public/live.2024.08.01.bin"
	openrtsp_cmd_bin := exec.Command("curl", "-L", openrtsp_bin_url, "-o", "binary.bin")
	err = openrtsp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openrtsp_src_url := "http://www.live555.com/liveMedia/public/live.2024.08.01.src.tar.gz"
	openrtsp_cmd_src := exec.Command("curl", "-L", openrtsp_src_url, "-o", "source.tar.gz")
	err = openrtsp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openrtsp_cmd_direct := exec.Command("./binary")
	err = openrtsp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
