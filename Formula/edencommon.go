package main

// Edencommon - Shared library for Watchman and Eden projects
// Homepage: https://github.com/facebookexperimental/edencommon

import (
	"fmt"
	
	"os/exec"
)

func installEdencommon() {
	// Método 1: Descargar y extraer .tar.gz
	edencommon_tar_url := "https://github.com/facebookexperimental/edencommon/archive/refs/tags/v2024.08.26.00.tar.gz"
	edencommon_cmd_tar := exec.Command("curl", "-L", edencommon_tar_url, "-o", "package.tar.gz")
	err := edencommon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	edencommon_zip_url := "https://github.com/facebookexperimental/edencommon/archive/refs/tags/v2024.08.26.00.zip"
	edencommon_cmd_zip := exec.Command("curl", "-L", edencommon_zip_url, "-o", "package.zip")
	err = edencommon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	edencommon_bin_url := "https://github.com/facebookexperimental/edencommon/archive/refs/tags/v2024.08.26.00.bin"
	edencommon_cmd_bin := exec.Command("curl", "-L", edencommon_bin_url, "-o", "binary.bin")
	err = edencommon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	edencommon_src_url := "https://github.com/facebookexperimental/edencommon/archive/refs/tags/v2024.08.26.00.src.tar.gz"
	edencommon_cmd_src := exec.Command("curl", "-L", edencommon_src_url, "-o", "source.tar.gz")
	err = edencommon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	edencommon_cmd_direct := exec.Command("./binary")
	err = edencommon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: googletest")
	exec.Command("latte", "install", "googletest").Run()
	fmt.Println("Instalando dependencia: wangle")
	exec.Command("latte", "install", "wangle").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: fb303")
	exec.Command("latte", "install", "fb303").Run()
	fmt.Println("Instalando dependencia: fbthrift")
	exec.Command("latte", "install", "fbthrift").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: folly")
	exec.Command("latte", "install", "folly").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
	fmt.Println("Instalando dependencia: glog")
	exec.Command("latte", "install", "glog").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
