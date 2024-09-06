package main

// Wbox - HTTP testing tool and configuration-less HTTP server
// Homepage: https://web.archive.org/web/20221105011338/http://www.hping.org/wbox/

import (
	"fmt"
	
	"os/exec"
)

func installWbox() {
	// Método 1: Descargar y extraer .tar.gz
	wbox_tar_url := "https://web.archive.org/web/20220524011612/http://www.hping.org/wbox/wbox-5.tar.gz"
	wbox_cmd_tar := exec.Command("curl", "-L", wbox_tar_url, "-o", "package.tar.gz")
	err := wbox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wbox_zip_url := "https://web.archive.org/web/20220524011612/http://www.hping.org/wbox/wbox-5.zip"
	wbox_cmd_zip := exec.Command("curl", "-L", wbox_zip_url, "-o", "package.zip")
	err = wbox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wbox_bin_url := "https://web.archive.org/web/20220524011612/http://www.hping.org/wbox/wbox-5.bin"
	wbox_cmd_bin := exec.Command("curl", "-L", wbox_bin_url, "-o", "binary.bin")
	err = wbox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wbox_src_url := "https://web.archive.org/web/20220524011612/http://www.hping.org/wbox/wbox-5.src.tar.gz"
	wbox_cmd_src := exec.Command("curl", "-L", wbox_src_url, "-o", "source.tar.gz")
	err = wbox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wbox_cmd_direct := exec.Command("./binary")
	err = wbox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
