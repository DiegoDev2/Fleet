package main

// TremorRuntime - Early-stage event processing system for unstructured data
// Homepage: https://www.tremor.rs/

import (
	"fmt"
	
	"os/exec"
)

func installTremorRuntime() {
	// Método 1: Descargar y extraer .tar.gz
	tremorruntime_tar_url := "https://github.com/tremor-rs/tremor-runtime/archive/refs/tags/v0.12.4.tar.gz"
	tremorruntime_cmd_tar := exec.Command("curl", "-L", tremorruntime_tar_url, "-o", "package.tar.gz")
	err := tremorruntime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tremorruntime_zip_url := "https://github.com/tremor-rs/tremor-runtime/archive/refs/tags/v0.12.4.zip"
	tremorruntime_cmd_zip := exec.Command("curl", "-L", tremorruntime_zip_url, "-o", "package.zip")
	err = tremorruntime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tremorruntime_bin_url := "https://github.com/tremor-rs/tremor-runtime/archive/refs/tags/v0.12.4.bin"
	tremorruntime_cmd_bin := exec.Command("curl", "-L", tremorruntime_bin_url, "-o", "binary.bin")
	err = tremorruntime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tremorruntime_src_url := "https://github.com/tremor-rs/tremor-runtime/archive/refs/tags/v0.12.4.src.tar.gz"
	tremorruntime_cmd_src := exec.Command("curl", "-L", tremorruntime_src_url, "-o", "source.tar.gz")
	err = tremorruntime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tremorruntime_cmd_direct := exec.Command("./binary")
	err = tremorruntime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: librdkafka")
	exec.Command("latte", "install", "librdkafka").Run()
	fmt.Println("Instalando dependencia: oniguruma")
	exec.Command("latte", "install", "oniguruma").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: llvm@15")
	exec.Command("latte", "install", "llvm@15").Run()
}
