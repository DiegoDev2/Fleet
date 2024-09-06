package main

// IosWebkitDebugProxy - DevTools proxy for iOS devices
// Homepage: https://github.com/google/ios-webkit-debug-proxy

import (
	"fmt"
	
	"os/exec"
)

func installIosWebkitDebugProxy() {
	// Método 1: Descargar y extraer .tar.gz
	ioswebkitdebugproxy_tar_url := "https://github.com/google/ios-webkit-debug-proxy/archive/refs/tags/v1.9.1.tar.gz"
	ioswebkitdebugproxy_cmd_tar := exec.Command("curl", "-L", ioswebkitdebugproxy_tar_url, "-o", "package.tar.gz")
	err := ioswebkitdebugproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ioswebkitdebugproxy_zip_url := "https://github.com/google/ios-webkit-debug-proxy/archive/refs/tags/v1.9.1.zip"
	ioswebkitdebugproxy_cmd_zip := exec.Command("curl", "-L", ioswebkitdebugproxy_zip_url, "-o", "package.zip")
	err = ioswebkitdebugproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ioswebkitdebugproxy_bin_url := "https://github.com/google/ios-webkit-debug-proxy/archive/refs/tags/v1.9.1.bin"
	ioswebkitdebugproxy_cmd_bin := exec.Command("curl", "-L", ioswebkitdebugproxy_bin_url, "-o", "binary.bin")
	err = ioswebkitdebugproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ioswebkitdebugproxy_src_url := "https://github.com/google/ios-webkit-debug-proxy/archive/refs/tags/v1.9.1.src.tar.gz"
	ioswebkitdebugproxy_cmd_src := exec.Command("curl", "-L", ioswebkitdebugproxy_src_url, "-o", "source.tar.gz")
	err = ioswebkitdebugproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ioswebkitdebugproxy_cmd_direct := exec.Command("./binary")
	err = ioswebkitdebugproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libimobiledevice")
	exec.Command("latte", "install", "libimobiledevice").Run()
	fmt.Println("Instalando dependencia: libplist")
	exec.Command("latte", "install", "libplist").Run()
	fmt.Println("Instalando dependencia: libusbmuxd")
	exec.Command("latte", "install", "libusbmuxd").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
