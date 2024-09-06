package main

// Gcsfuse - User-space file system for interacting with Google Cloud
// Homepage: https://github.com/googlecloudplatform/gcsfuse

import (
	"fmt"
	
	"os/exec"
)

func installGcsfuse() {
	// Método 1: Descargar y extraer .tar.gz
	gcsfuse_tar_url := "https://github.com/GoogleCloudPlatform/gcsfuse/archive/refs/tags/v2.4.0.tar.gz"
	gcsfuse_cmd_tar := exec.Command("curl", "-L", gcsfuse_tar_url, "-o", "package.tar.gz")
	err := gcsfuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gcsfuse_zip_url := "https://github.com/GoogleCloudPlatform/gcsfuse/archive/refs/tags/v2.4.0.zip"
	gcsfuse_cmd_zip := exec.Command("curl", "-L", gcsfuse_zip_url, "-o", "package.zip")
	err = gcsfuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gcsfuse_bin_url := "https://github.com/GoogleCloudPlatform/gcsfuse/archive/refs/tags/v2.4.0.bin"
	gcsfuse_cmd_bin := exec.Command("curl", "-L", gcsfuse_bin_url, "-o", "binary.bin")
	err = gcsfuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gcsfuse_src_url := "https://github.com/GoogleCloudPlatform/gcsfuse/archive/refs/tags/v2.4.0.src.tar.gz"
	gcsfuse_cmd_src := exec.Command("curl", "-L", gcsfuse_src_url, "-o", "source.tar.gz")
	err = gcsfuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gcsfuse_cmd_direct := exec.Command("./binary")
	err = gcsfuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: libfuse")
	exec.Command("latte", "install", "libfuse").Run()
}
