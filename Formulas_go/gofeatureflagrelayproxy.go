package main

// GoFeatureFlagRelayProxy - Stand alone server to run GO Feature Flag
// Homepage: https://gofeatureflag.org

import (
	"fmt"
	
	"os/exec"
)

func installGoFeatureFlagRelayProxy() {
	// Método 1: Descargar y extraer .tar.gz
	gofeatureflagrelayproxy_tar_url := "https://github.com/thomaspoignant/go-feature-flag/archive/refs/tags/v1.33.0.tar.gz"
	gofeatureflagrelayproxy_cmd_tar := exec.Command("curl", "-L", gofeatureflagrelayproxy_tar_url, "-o", "package.tar.gz")
	err := gofeatureflagrelayproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gofeatureflagrelayproxy_zip_url := "https://github.com/thomaspoignant/go-feature-flag/archive/refs/tags/v1.33.0.zip"
	gofeatureflagrelayproxy_cmd_zip := exec.Command("curl", "-L", gofeatureflagrelayproxy_zip_url, "-o", "package.zip")
	err = gofeatureflagrelayproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gofeatureflagrelayproxy_bin_url := "https://github.com/thomaspoignant/go-feature-flag/archive/refs/tags/v1.33.0.bin"
	gofeatureflagrelayproxy_cmd_bin := exec.Command("curl", "-L", gofeatureflagrelayproxy_bin_url, "-o", "binary.bin")
	err = gofeatureflagrelayproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gofeatureflagrelayproxy_src_url := "https://github.com/thomaspoignant/go-feature-flag/archive/refs/tags/v1.33.0.src.tar.gz"
	gofeatureflagrelayproxy_cmd_src := exec.Command("curl", "-L", gofeatureflagrelayproxy_src_url, "-o", "source.tar.gz")
	err = gofeatureflagrelayproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gofeatureflagrelayproxy_cmd_direct := exec.Command("./binary")
	err = gofeatureflagrelayproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
