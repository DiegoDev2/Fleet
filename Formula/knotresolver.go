package main

// KnotResolver - Minimalistic, caching, DNSSEC-validating DNS resolver
// Homepage: https://www.knot-resolver.cz

import (
	"fmt"
	
	"os/exec"
)

func installKnotResolver() {
	// Método 1: Descargar y extraer .tar.gz
	knotresolver_tar_url := "https://secure.nic.cz/files/knot-resolver/knot-resolver-5.7.2.tar.xz"
	knotresolver_cmd_tar := exec.Command("curl", "-L", knotresolver_tar_url, "-o", "package.tar.gz")
	err := knotresolver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	knotresolver_zip_url := "https://secure.nic.cz/files/knot-resolver/knot-resolver-5.7.2.tar.xz"
	knotresolver_cmd_zip := exec.Command("curl", "-L", knotresolver_zip_url, "-o", "package.zip")
	err = knotresolver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	knotresolver_bin_url := "https://secure.nic.cz/files/knot-resolver/knot-resolver-5.7.2.tar.xz"
	knotresolver_cmd_bin := exec.Command("curl", "-L", knotresolver_bin_url, "-o", "binary.bin")
	err = knotresolver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	knotresolver_src_url := "https://secure.nic.cz/files/knot-resolver/knot-resolver-5.7.2.tar.xz"
	knotresolver_cmd_src := exec.Command("curl", "-L", knotresolver_src_url, "-o", "source.tar.gz")
	err = knotresolver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	knotresolver_cmd_direct := exec.Command("./binary")
	err = knotresolver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: knot")
	exec.Command("latte", "install", "knot").Run()
	fmt.Println("Instalando dependencia: libnghttp2")
	exec.Command("latte", "install", "libnghttp2").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: lmdb")
	exec.Command("latte", "install", "lmdb").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
	fmt.Println("Instalando dependencia: libcap-ng")
	exec.Command("latte", "install", "libcap-ng").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
