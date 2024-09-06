package main

// Libnatpmp - NAT port mapping protocol library
// Homepage: http://miniupnp.free.fr/libnatpmp.html

import (
	"fmt"
	
	"os/exec"
)

func installLibnatpmp() {
	// Método 1: Descargar y extraer .tar.gz
	libnatpmp_tar_url := "http://miniupnp.free.fr/files/download.php?file=libnatpmp-20230423.tar.gz"
	libnatpmp_cmd_tar := exec.Command("curl", "-L", libnatpmp_tar_url, "-o", "package.tar.gz")
	err := libnatpmp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnatpmp_zip_url := "http://miniupnp.free.fr/files/download.php?file=libnatpmp-20230423.zip"
	libnatpmp_cmd_zip := exec.Command("curl", "-L", libnatpmp_zip_url, "-o", "package.zip")
	err = libnatpmp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnatpmp_bin_url := "http://miniupnp.free.fr/files/download.php?file=libnatpmp-20230423.bin"
	libnatpmp_cmd_bin := exec.Command("curl", "-L", libnatpmp_bin_url, "-o", "binary.bin")
	err = libnatpmp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnatpmp_src_url := "http://miniupnp.free.fr/files/download.php?file=libnatpmp-20230423.src.tar.gz"
	libnatpmp_cmd_src := exec.Command("curl", "-L", libnatpmp_src_url, "-o", "source.tar.gz")
	err = libnatpmp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnatpmp_cmd_direct := exec.Command("./binary")
	err = libnatpmp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
