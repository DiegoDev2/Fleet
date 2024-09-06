package main

// Samba - SMB/CIFS file, print, and login server for UNIX
// Homepage: https://www.samba.org/

import (
	"fmt"
	
	"os/exec"
)

func installSamba() {
	// Método 1: Descargar y extraer .tar.gz
	samba_tar_url := "https://download.samba.org/pub/samba/stable/samba-4.21.0.tar.gz"
	samba_cmd_tar := exec.Command("curl", "-L", samba_tar_url, "-o", "package.tar.gz")
	err := samba_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	samba_zip_url := "https://download.samba.org/pub/samba/stable/samba-4.21.0.zip"
	samba_cmd_zip := exec.Command("curl", "-L", samba_zip_url, "-o", "package.zip")
	err = samba_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	samba_bin_url := "https://download.samba.org/pub/samba/stable/samba-4.21.0.bin"
	samba_cmd_bin := exec.Command("curl", "-L", samba_bin_url, "-o", "binary.bin")
	err = samba_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	samba_src_url := "https://download.samba.org/pub/samba/stable/samba-4.21.0.src.tar.gz"
	samba_cmd_src := exec.Command("curl", "-L", samba_src_url, "-o", "source.tar.gz")
	err = samba_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	samba_cmd_direct := exec.Command("./binary")
	err = samba_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmocka")
exec.Command("latte", "install", "cmocka")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: krb5")
exec.Command("latte", "install", "krb5")
	fmt.Println("Instalando dependencia: libtasn1")
exec.Command("latte", "install", "libtasn1")
	fmt.Println("Instalando dependencia: popt")
exec.Command("latte", "install", "popt")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: talloc")
exec.Command("latte", "install", "talloc")
	fmt.Println("Instalando dependencia: tdb")
exec.Command("latte", "install", "tdb")
	fmt.Println("Instalando dependencia: tevent")
exec.Command("latte", "install", "tevent")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
}
