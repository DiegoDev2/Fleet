package main

// CoreLightning - Lightning Network implementation focusing on spec compliance and performance
// Homepage: https://github.com/ElementsProject/lightning

import (
	"fmt"
	
	"os/exec"
)

func installCoreLightning() {
	// Método 1: Descargar y extraer .tar.gz
	corelightning_tar_url := "https://github.com/ElementsProject/lightning/releases/download/v24.05/clightning-v24.05.zip"
	corelightning_cmd_tar := exec.Command("curl", "-L", corelightning_tar_url, "-o", "package.tar.gz")
	err := corelightning_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	corelightning_zip_url := "https://github.com/ElementsProject/lightning/releases/download/v24.05/clightning-v24.05.zip"
	corelightning_cmd_zip := exec.Command("curl", "-L", corelightning_zip_url, "-o", "package.zip")
	err = corelightning_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	corelightning_bin_url := "https://github.com/ElementsProject/lightning/releases/download/v24.05/clightning-v24.05.zip"
	corelightning_cmd_bin := exec.Command("curl", "-L", corelightning_bin_url, "-o", "binary.bin")
	err = corelightning_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	corelightning_src_url := "https://github.com/ElementsProject/lightning/releases/download/v24.05/clightning-v24.05.zip"
	corelightning_cmd_src := exec.Command("curl", "-L", corelightning_src_url, "-o", "source.tar.gz")
	err = corelightning_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	corelightning_cmd_direct := exec.Command("./binary")
	err = corelightning_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: jq")
exec.Command("latte", "install", "jq")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: lowdown")
exec.Command("latte", "install", "lowdown")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: poetry")
exec.Command("latte", "install", "poetry")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: bitcoin")
exec.Command("latte", "install", "bitcoin")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
}
