package main

// Libhdhomerun - C library for controlling SiliconDust HDHomeRun TV tuners
// Homepage: https://www.silicondust.com/support/linux/

import (
	"fmt"
	
	"os/exec"
)

func installLibhdhomerun() {
	// Método 1: Descargar y extraer .tar.gz
	libhdhomerun_tar_url := "https://download.silicondust.com/hdhomerun/libhdhomerun_20231214.tgz"
	libhdhomerun_cmd_tar := exec.Command("curl", "-L", libhdhomerun_tar_url, "-o", "package.tar.gz")
	err := libhdhomerun_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libhdhomerun_zip_url := "https://download.silicondust.com/hdhomerun/libhdhomerun_20231214.tgz"
	libhdhomerun_cmd_zip := exec.Command("curl", "-L", libhdhomerun_zip_url, "-o", "package.zip")
	err = libhdhomerun_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libhdhomerun_bin_url := "https://download.silicondust.com/hdhomerun/libhdhomerun_20231214.tgz"
	libhdhomerun_cmd_bin := exec.Command("curl", "-L", libhdhomerun_bin_url, "-o", "binary.bin")
	err = libhdhomerun_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libhdhomerun_src_url := "https://download.silicondust.com/hdhomerun/libhdhomerun_20231214.tgz"
	libhdhomerun_cmd_src := exec.Command("curl", "-L", libhdhomerun_src_url, "-o", "source.tar.gz")
	err = libhdhomerun_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libhdhomerun_cmd_direct := exec.Command("./binary")
	err = libhdhomerun_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
