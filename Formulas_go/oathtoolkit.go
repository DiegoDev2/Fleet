package main

// OathToolkit - Tools for one-time password authentication systems
// Homepage: https://www.nongnu.org/oath-toolkit/

import (
	"fmt"
	
	"os/exec"
)

func installOathToolkit() {
	// Método 1: Descargar y extraer .tar.gz
	oathtoolkit_tar_url := "https://download-mirror.savannah.gnu.org/releases/oath-toolkit/oath-toolkit-2.6.11.tar.gz"
	oathtoolkit_cmd_tar := exec.Command("curl", "-L", oathtoolkit_tar_url, "-o", "package.tar.gz")
	err := oathtoolkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oathtoolkit_zip_url := "https://download-mirror.savannah.gnu.org/releases/oath-toolkit/oath-toolkit-2.6.11.zip"
	oathtoolkit_cmd_zip := exec.Command("curl", "-L", oathtoolkit_zip_url, "-o", "package.zip")
	err = oathtoolkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oathtoolkit_bin_url := "https://download-mirror.savannah.gnu.org/releases/oath-toolkit/oath-toolkit-2.6.11.bin"
	oathtoolkit_cmd_bin := exec.Command("curl", "-L", oathtoolkit_bin_url, "-o", "binary.bin")
	err = oathtoolkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oathtoolkit_src_url := "https://download-mirror.savannah.gnu.org/releases/oath-toolkit/oath-toolkit-2.6.11.src.tar.gz"
	oathtoolkit_cmd_src := exec.Command("curl", "-L", oathtoolkit_src_url, "-o", "source.tar.gz")
	err = oathtoolkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oathtoolkit_cmd_direct := exec.Command("./binary")
	err = oathtoolkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gtk-doc")
exec.Command("latte", "install", "gtk-doc")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: libxmlsec1")
exec.Command("latte", "install", "libxmlsec1")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
