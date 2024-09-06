package main

// Apt - Advanced Package Tool
// Homepage: https://wiki.debian.org/Apt

import (
	"fmt"
	
	"os/exec"
)

func installApt() {
	// Método 1: Descargar y extraer .tar.gz
	apt_tar_url := "https://deb.debian.org/debian/pool/main/a/apt/apt_2.9.8.tar.xz"
	apt_cmd_tar := exec.Command("curl", "-L", apt_tar_url, "-o", "package.tar.gz")
	err := apt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apt_zip_url := "https://deb.debian.org/debian/pool/main/a/apt/apt_2.9.8.tar.xz"
	apt_cmd_zip := exec.Command("curl", "-L", apt_zip_url, "-o", "package.zip")
	err = apt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apt_bin_url := "https://deb.debian.org/debian/pool/main/a/apt/apt_2.9.8.tar.xz"
	apt_cmd_bin := exec.Command("curl", "-L", apt_bin_url, "-o", "binary.bin")
	err = apt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apt_src_url := "https://deb.debian.org/debian/pool/main/a/apt/apt_2.9.8.tar.xz"
	apt_cmd_src := exec.Command("curl", "-L", apt_src_url, "-o", "source.tar.gz")
	err = apt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apt_cmd_direct := exec.Command("./binary")
	err = apt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: docbook")
exec.Command("latte", "install", "docbook")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: googletest")
exec.Command("latte", "install", "googletest")
	fmt.Println("Instalando dependencia: libxslt")
exec.Command("latte", "install", "libxslt")
	fmt.Println("Instalando dependencia: po4a")
exec.Command("latte", "install", "po4a")
	fmt.Println("Instalando dependencia: w3m")
exec.Command("latte", "install", "w3m")
	fmt.Println("Instalando dependencia: berkeley-db@5")
exec.Command("latte", "install", "berkeley-db@5")
	fmt.Println("Instalando dependencia: bzip2")
exec.Command("latte", "install", "bzip2")
	fmt.Println("Instalando dependencia: dpkg")
exec.Command("latte", "install", "dpkg")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gnupg")
exec.Command("latte", "install", "gnupg")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
	fmt.Println("Instalando dependencia: xxhash")
exec.Command("latte", "install", "xxhash")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
