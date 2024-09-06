package main

// ReactNativeCli - Tools for creating native apps for Android and iOS
// Homepage: https://facebook.github.io/react-native/

import (
	"fmt"
	
	"os/exec"
)

func installReactNativeCli() {
	// Método 1: Descargar y extraer .tar.gz
	reactnativecli_tar_url := "https://registry.npmjs.org/react-native-cli/-/react-native-cli-2.0.1.tgz"
	reactnativecli_cmd_tar := exec.Command("curl", "-L", reactnativecli_tar_url, "-o", "package.tar.gz")
	err := reactnativecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reactnativecli_zip_url := "https://registry.npmjs.org/react-native-cli/-/react-native-cli-2.0.1.tgz"
	reactnativecli_cmd_zip := exec.Command("curl", "-L", reactnativecli_zip_url, "-o", "package.zip")
	err = reactnativecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reactnativecli_bin_url := "https://registry.npmjs.org/react-native-cli/-/react-native-cli-2.0.1.tgz"
	reactnativecli_cmd_bin := exec.Command("curl", "-L", reactnativecli_bin_url, "-o", "binary.bin")
	err = reactnativecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reactnativecli_src_url := "https://registry.npmjs.org/react-native-cli/-/react-native-cli-2.0.1.tgz"
	reactnativecli_cmd_src := exec.Command("curl", "-L", reactnativecli_src_url, "-o", "source.tar.gz")
	err = reactnativecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reactnativecli_cmd_direct := exec.Command("./binary")
	err = reactnativecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
