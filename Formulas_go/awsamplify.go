package main

// AwsAmplify - Build full-stack web and mobile apps in hours. Easy to start, easy to scale
// Homepage: https://aws.amazon.com/amplify

import (
	"fmt"
	
	"os/exec"
)

func installAwsAmplify() {
	// Método 1: Descargar y extraer .tar.gz
	awsamplify_tar_url := "https://registry.npmjs.org/@aws-amplify/cli-internal/-/cli-internal-12.12.6.tgz"
	awsamplify_cmd_tar := exec.Command("curl", "-L", awsamplify_tar_url, "-o", "package.tar.gz")
	err := awsamplify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsamplify_zip_url := "https://registry.npmjs.org/@aws-amplify/cli-internal/-/cli-internal-12.12.6.tgz"
	awsamplify_cmd_zip := exec.Command("curl", "-L", awsamplify_zip_url, "-o", "package.zip")
	err = awsamplify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsamplify_bin_url := "https://registry.npmjs.org/@aws-amplify/cli-internal/-/cli-internal-12.12.6.tgz"
	awsamplify_cmd_bin := exec.Command("curl", "-L", awsamplify_bin_url, "-o", "binary.bin")
	err = awsamplify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsamplify_src_url := "https://registry.npmjs.org/@aws-amplify/cli-internal/-/cli-internal-12.12.6.tgz"
	awsamplify_cmd_src := exec.Command("curl", "-L", awsamplify_src_url, "-o", "source.tar.gz")
	err = awsamplify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsamplify_cmd_direct := exec.Command("./binary")
	err = awsamplify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
