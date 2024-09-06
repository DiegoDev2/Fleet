package main

// NicovideoDl - Command-line program to download videos from www.nicovideo.jp
// Homepage: https://osdn.net/projects/nicovideo-dl/

import (
	"fmt"
	
	"os/exec"
)

func installNicovideoDl() {
	// Método 1: Descargar y extraer .tar.gz
	nicovideodl_tar_url := "https://dotsrc.dl.osdn.net/osdn/nicovideo-dl/70568/nicovideo-dl-0.0.20190126.tar.gz"
	nicovideodl_cmd_tar := exec.Command("curl", "-L", nicovideodl_tar_url, "-o", "package.tar.gz")
	err := nicovideodl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nicovideodl_zip_url := "https://dotsrc.dl.osdn.net/osdn/nicovideo-dl/70568/nicovideo-dl-0.0.20190126.zip"
	nicovideodl_cmd_zip := exec.Command("curl", "-L", nicovideodl_zip_url, "-o", "package.zip")
	err = nicovideodl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nicovideodl_bin_url := "https://dotsrc.dl.osdn.net/osdn/nicovideo-dl/70568/nicovideo-dl-0.0.20190126.bin"
	nicovideodl_cmd_bin := exec.Command("curl", "-L", nicovideodl_bin_url, "-o", "binary.bin")
	err = nicovideodl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nicovideodl_src_url := "https://dotsrc.dl.osdn.net/osdn/nicovideo-dl/70568/nicovideo-dl-0.0.20190126.src.tar.gz"
	nicovideodl_cmd_src := exec.Command("curl", "-L", nicovideodl_src_url, "-o", "source.tar.gz")
	err = nicovideodl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nicovideodl_cmd_direct := exec.Command("./binary")
	err = nicovideodl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
