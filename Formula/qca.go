package main

// Qca - Qt Cryptographic Architecture (QCA)
// Homepage: https://userbase.kde.org/QCA

import (
	"fmt"
	
	"os/exec"
)

func installQca() {
	// Método 1: Descargar y extraer .tar.gz
	qca_tar_url := "https://download.kde.org/stable/qca/2.3.9/qca-2.3.9.tar.xz"
	qca_cmd_tar := exec.Command("curl", "-L", qca_tar_url, "-o", "package.tar.gz")
	err := qca_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qca_zip_url := "https://download.kde.org/stable/qca/2.3.9/qca-2.3.9.tar.xz"
	qca_cmd_zip := exec.Command("curl", "-L", qca_zip_url, "-o", "package.zip")
	err = qca_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qca_bin_url := "https://download.kde.org/stable/qca/2.3.9/qca-2.3.9.tar.xz"
	qca_cmd_bin := exec.Command("curl", "-L", qca_bin_url, "-o", "binary.bin")
	err = qca_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qca_src_url := "https://download.kde.org/stable/qca/2.3.9/qca-2.3.9.tar.xz"
	qca_cmd_src := exec.Command("curl", "-L", qca_src_url, "-o", "source.tar.gz")
	err = qca_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qca_cmd_direct := exec.Command("./binary")
	err = qca_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: botan")
	exec.Command("latte", "install", "botan").Run()
	fmt.Println("Instalando dependencia: ca-certificates")
	exec.Command("latte", "install", "ca-certificates").Run()
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: nss")
	exec.Command("latte", "install", "nss").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkcs11-helper")
	exec.Command("latte", "install", "pkcs11-helper").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: nspr")
	exec.Command("latte", "install", "nspr").Run()
}
