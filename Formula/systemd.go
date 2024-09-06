package main

// Systemd - System and service manager
// Homepage: https://systemd.io

import (
	"fmt"
	
	"os/exec"
)

func installSystemd() {
	// Método 1: Descargar y extraer .tar.gz
	systemd_tar_url := "https://github.com/systemd/systemd/archive/refs/tags/v256.4.tar.gz"
	systemd_cmd_tar := exec.Command("curl", "-L", systemd_tar_url, "-o", "package.tar.gz")
	err := systemd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	systemd_zip_url := "https://github.com/systemd/systemd/archive/refs/tags/v256.4.zip"
	systemd_cmd_zip := exec.Command("curl", "-L", systemd_zip_url, "-o", "package.zip")
	err = systemd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	systemd_bin_url := "https://github.com/systemd/systemd/archive/refs/tags/v256.4.bin"
	systemd_cmd_bin := exec.Command("curl", "-L", systemd_bin_url, "-o", "binary.bin")
	err = systemd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	systemd_src_url := "https://github.com/systemd/systemd/archive/refs/tags/v256.4.src.tar.gz"
	systemd_cmd_src := exec.Command("curl", "-L", systemd_src_url, "-o", "source.tar.gz")
	err = systemd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	systemd_cmd_direct := exec.Command("./binary")
	err = systemd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gperf")
	exec.Command("latte", "install", "gperf").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: libxslt")
	exec.Command("latte", "install", "libxslt").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libcap")
	exec.Command("latte", "install", "libcap").Run()
	fmt.Println("Instalando dependencia: libxcrypt")
	exec.Command("latte", "install", "libxcrypt").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
