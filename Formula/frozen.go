package main

// Frozen - Header-only, constexpr alternative to gperf for C++14 users
// Homepage: https://github.com/serge-sans-paille/frozen

import (
	"fmt"
	
	"os/exec"
)

func installFrozen() {
	// Método 1: Descargar y extraer .tar.gz
	frozen_tar_url := "https://github.com/serge-sans-paille/frozen/archive/refs/tags/1.2.0.tar.gz"
	frozen_cmd_tar := exec.Command("curl", "-L", frozen_tar_url, "-o", "package.tar.gz")
	err := frozen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	frozen_zip_url := "https://github.com/serge-sans-paille/frozen/archive/refs/tags/1.2.0.zip"
	frozen_cmd_zip := exec.Command("curl", "-L", frozen_zip_url, "-o", "package.zip")
	err = frozen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	frozen_bin_url := "https://github.com/serge-sans-paille/frozen/archive/refs/tags/1.2.0.bin"
	frozen_cmd_bin := exec.Command("curl", "-L", frozen_bin_url, "-o", "binary.bin")
	err = frozen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	frozen_src_url := "https://github.com/serge-sans-paille/frozen/archive/refs/tags/1.2.0.src.tar.gz"
	frozen_cmd_src := exec.Command("curl", "-L", frozen_src_url, "-o", "source.tar.gz")
	err = frozen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	frozen_cmd_direct := exec.Command("./binary")
	err = frozen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
