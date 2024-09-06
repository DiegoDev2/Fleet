package main

// Rpm - Standard unix software packaging tool
// Homepage: https://rpm.org/

import (
	"fmt"
	
	"os/exec"
)

func installRpm() {
	// Método 1: Descargar y extraer .tar.gz
	rpm_tar_url := "https://ftp.osuosl.org/pub/rpm/releases/rpm-4.19.x/rpm-4.19.1.1.tar.bz2"
	rpm_cmd_tar := exec.Command("curl", "-L", rpm_tar_url, "-o", "package.tar.gz")
	err := rpm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rpm_zip_url := "https://ftp.osuosl.org/pub/rpm/releases/rpm-4.19.x/rpm-4.19.1.1.tar.bz2"
	rpm_cmd_zip := exec.Command("curl", "-L", rpm_zip_url, "-o", "package.zip")
	err = rpm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rpm_bin_url := "https://ftp.osuosl.org/pub/rpm/releases/rpm-4.19.x/rpm-4.19.1.1.tar.bz2"
	rpm_cmd_bin := exec.Command("curl", "-L", rpm_bin_url, "-o", "binary.bin")
	err = rpm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rpm_src_url := "https://ftp.osuosl.org/pub/rpm/releases/rpm-4.19.x/rpm-4.19.1.1.tar.bz2"
	rpm_cmd_src := exec.Command("curl", "-L", rpm_src_url, "-o", "source.tar.gz")
	err = rpm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rpm_cmd_direct := exec.Command("./binary")
	err = rpm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: gawk")
exec.Command("latte", "install", "gawk")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: popt")
exec.Command("latte", "install", "popt")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
