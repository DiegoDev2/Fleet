package main

// Php - General-purpose scripting language
// Homepage: https://www.php.net/

import (
	"fmt"
	
	"os/exec"
)

func installPhp() {
	// Método 1: Descargar y extraer .tar.gz
	php_tar_url := "https://www.php.net/distributions/php-8.3.11.tar.xz"
	php_cmd_tar := exec.Command("curl", "-L", php_tar_url, "-o", "package.tar.gz")
	err := php_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	php_zip_url := "https://www.php.net/distributions/php-8.3.11.tar.xz"
	php_cmd_zip := exec.Command("curl", "-L", php_zip_url, "-o", "package.zip")
	err = php_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	php_bin_url := "https://www.php.net/distributions/php-8.3.11.tar.xz"
	php_cmd_bin := exec.Command("curl", "-L", php_bin_url, "-o", "binary.bin")
	err = php_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	php_src_url := "https://www.php.net/distributions/php-8.3.11.tar.xz"
	php_cmd_src := exec.Command("curl", "-L", php_src_url, "-o", "source.tar.gz")
	err = php_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	php_cmd_direct := exec.Command("./binary")
	err = php_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: re2c")
	exec.Command("latte", "install", "re2c").Run()
	fmt.Println("Instalando dependencia: httpd")
	exec.Command("latte", "install", "httpd").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: apr")
	exec.Command("latte", "install", "apr").Run()
	fmt.Println("Instalando dependencia: apr-util")
	exec.Command("latte", "install", "apr-util").Run()
	fmt.Println("Instalando dependencia: argon2")
	exec.Command("latte", "install", "argon2").Run()
	fmt.Println("Instalando dependencia: aspell")
	exec.Command("latte", "install", "aspell").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: freetds")
	exec.Command("latte", "install", "freetds").Run()
	fmt.Println("Instalando dependencia: gd")
	exec.Command("latte", "install", "gd").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: krb5")
	exec.Command("latte", "install", "krb5").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
	fmt.Println("Instalando dependencia: oniguruma")
	exec.Command("latte", "install", "oniguruma").Run()
	fmt.Println("Instalando dependencia: openldap")
	exec.Command("latte", "install", "openldap").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: tidy-html5")
	exec.Command("latte", "install", "tidy-html5").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
}
