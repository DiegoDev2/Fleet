package main

// Libmodplug - Library from the Modplug-XMMS project
// Homepage: https://modplug-xmms.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibmodplug() {
	// Método 1: Descargar y extraer .tar.gz
	libmodplug_tar_url := "https://downloads.sourceforge.net/project/modplug-xmms/libmodplug/0.8.9.0/libmodplug-0.8.9.0.tar.gz"
	libmodplug_cmd_tar := exec.Command("curl", "-L", libmodplug_tar_url, "-o", "package.tar.gz")
	err := libmodplug_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmodplug_zip_url := "https://downloads.sourceforge.net/project/modplug-xmms/libmodplug/0.8.9.0/libmodplug-0.8.9.0.zip"
	libmodplug_cmd_zip := exec.Command("curl", "-L", libmodplug_zip_url, "-o", "package.zip")
	err = libmodplug_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmodplug_bin_url := "https://downloads.sourceforge.net/project/modplug-xmms/libmodplug/0.8.9.0/libmodplug-0.8.9.0.bin"
	libmodplug_cmd_bin := exec.Command("curl", "-L", libmodplug_bin_url, "-o", "binary.bin")
	err = libmodplug_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmodplug_src_url := "https://downloads.sourceforge.net/project/modplug-xmms/libmodplug/0.8.9.0/libmodplug-0.8.9.0.src.tar.gz"
	libmodplug_cmd_src := exec.Command("curl", "-L", libmodplug_src_url, "-o", "source.tar.gz")
	err = libmodplug_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmodplug_cmd_direct := exec.Command("./binary")
	err = libmodplug_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
