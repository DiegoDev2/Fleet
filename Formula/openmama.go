package main

// Openmama - Open source high performance messaging API for various Market Data sources
// Homepage: https://openmama.finos.org

import (
	"fmt"
	
	"os/exec"
)

func installOpenmama() {
	// Método 1: Descargar y extraer .tar.gz
	openmama_tar_url := "https://github.com/finos/OpenMAMA/archive/refs/tags/OpenMAMA-6.3.2-release.tar.gz"
	openmama_cmd_tar := exec.Command("curl", "-L", openmama_tar_url, "-o", "package.tar.gz")
	err := openmama_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openmama_zip_url := "https://github.com/finos/OpenMAMA/archive/refs/tags/OpenMAMA-6.3.2-release.zip"
	openmama_cmd_zip := exec.Command("curl", "-L", openmama_zip_url, "-o", "package.zip")
	err = openmama_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openmama_bin_url := "https://github.com/finos/OpenMAMA/archive/refs/tags/OpenMAMA-6.3.2-release.bin"
	openmama_cmd_bin := exec.Command("curl", "-L", openmama_bin_url, "-o", "binary.bin")
	err = openmama_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openmama_src_url := "https://github.com/finos/OpenMAMA/archive/refs/tags/OpenMAMA-6.3.2-release.src.tar.gz"
	openmama_cmd_src := exec.Command("curl", "-L", openmama_src_url, "-o", "source.tar.gz")
	err = openmama_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openmama_cmd_direct := exec.Command("./binary")
	err = openmama_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: apr")
	exec.Command("latte", "install", "apr").Run()
	fmt.Println("Instalando dependencia: apr-util")
	exec.Command("latte", "install", "apr-util").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: qpid-proton")
	exec.Command("latte", "install", "qpid-proton").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
